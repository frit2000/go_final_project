package httpServer

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func getTask(w http.ResponseWriter, r *http.Request) {
	var tasks = map[string][]Task{
		"tasks": {},
	}
	var task Task

	// подключаемся к БД
	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		log.Println("ошибка при подключении к БД:", err)
	}
	defer db.Close()

	//запрос в базу на получение данных из таблицы
	rows, err := db.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date")
	if err != nil {
		log.Println("Ошибка запроса в базу:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		tasks["tasks"] = append(tasks["tasks"], task)
	}
	if err != nil {
		log.Println("Ошибка чтения рядов после запроса в базу")
		return
	}

	resp, err := json.Marshal(&tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
