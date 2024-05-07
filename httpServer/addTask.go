package httpServer

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/frit2000/go_final_project/nextdate"
)

func addTask(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var task Task
	var respTaskAdd RespTaskAdd

	// получаем данные из веб-интерфейса
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//переводим данные в стркутуру task
	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// проверяем что все поля date и title в task валидные
	err = checkFieldsTask(&task)
	if err != nil {
		respTaskAdd.Err = err
	}
	// подключаемся к БД
	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		log.Println("ошибка при подключении к БД:", err)
	}
	defer db.Close()

	//записываем поля структуры task в БД
	res, err := db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err != nil {
		log.Println("ошибка при добавлении записи БД:", err)
	}

	//получаем ID последней добавленной записи
	respTaskAdd.Id, err = res.LastInsertId()
	if err != nil {
		log.Println("ошибка получении последнего ID:", err)
	}

	resp, err := json.Marshal(&respTaskAdd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func checkFieldsTask(task *Task) error {
	if task.Title == "" {
		return fmt.Errorf("не указан заголовок задачи")
	}

	if task.Date == "" {
		task.Date = time.Now().Format("20060102")
		return nil
	}
	_, err := time.Parse("20060102", task.Date)
	if err != nil {
		return fmt.Errorf("дата неверного формата")
	}

	newDate := time.Now().Format("20060102")
	err = nil
	if task.Repeat != "" {
		newDate, err = nextdate.NextDate(time.Now(), task.Date, task.Repeat)
	}

	if task.Date < time.Now().Format("20060102") {
		task.Date = newDate
	}

	return err
}
