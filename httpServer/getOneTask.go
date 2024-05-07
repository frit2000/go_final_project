package httpServer

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func getOneTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	var respTask RespTaskError
	var resp []byte
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
		respTask.Err = "Задача не найдена"
	}

	if respTask.Err == "" {
		resp, err = json.Marshal(&task)
	} else {
		resp, err = json.Marshal(&respTask)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
