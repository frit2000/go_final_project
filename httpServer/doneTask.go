package httpServer

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/frit2000/go_final_project/nextdate"
)

func doneTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	var respTask RespTaskError

	id := r.FormValue("id")

	// подключаемся к БД
	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		log.Println("ошибка при подключении к БД:", err)
	}
	defer db.Close()

	err = db.QueryRow("SELECT * FROM scheduler WHERE id = :id", sql.Named("id", id)).Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		log.Println("ошибка чтении данных по id:", err)
	}

	//проверяем, есть ли такой ID задачи
	if len(task.Id) == 0 {
		respTask.Err = "Ошибка, нет такого ID"
	}

	if task.Repeat == "" {
		_, err = db.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
		if err != nil {
			log.Println("ошибка при обновлении записи БД:", err)
		}
	} else {
		newDate, err := nextdate.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			log.Println("ошибка при вычислении новой даты", err)
		}
		_, err = db.Exec("UPDATE scheduler SET date = :date WHERE id = :id",
			sql.Named("date", newDate),
			sql.Named("id", task.Id))
		if err != nil {
			log.Println("ошибка при обновлении записи БД:", err)
		}
	}

	resp, err := json.Marshal(&respTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
