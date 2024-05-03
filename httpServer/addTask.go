package httpServer

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Task struct {
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"note"`
	Repeat  string `json:"repeat,omitempty"`
}

// type addTaskResult struct {
// 	Id  []byte `json:"id"`
// 	Err error  `json:"error"`
// }

func addTask(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var task Task
	//	var taskResult addTaskResult

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		log.Println("ошибка при подключении к БД:", err)
	}
	defer db.Close()

	res, err := db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err != nil {
		log.Println("ошибка при добавлении записи БД:", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("ошибка получении последнего ID:", err)
	}

	resp, err := json.Marshal(&id) //или ошибку, потом посмотреть
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
