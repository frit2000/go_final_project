package httpServer

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func getTask(w http.ResponseWriter, r *http.Request) {
	var tasks = map[string][]Task{
		"tasks": {},
	}
	var task Task
	var rows *sql.Rows
	// подключаемся к БД
	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		log.Println("ошибка при подключении к БД:", err)
	}
	defer db.Close()

	//если нажат поиск, то выбираем записи согласно строке поиска
	searchString := r.FormValue("search")
	if searchString != "" {
		searchDate, errParse := time.Parse("02.01.2006", searchString)
		//если в поиске дата
		if errParse == nil {
			rows, err = db.Query("SELECT * FROM scheduler WHERE date = :searchString",
				sql.Named("searchString", searchDate.Format("20060102")),
				sql.Named("limit", 15))
			//если в поиске НЕ дата
		} else {
			rows, err = db.Query("SELECT * FROM scheduler WHERE title LIKE :searchString OR comment LIKE :searchString ORDER BY date LIMIT :limit",
				sql.Named("searchString", "%"+searchString+"%"),
				sql.Named("limit", 15))
		}

		//если НЕ нажат поиск, то выбираем все записи
	} else {
		rows, err = db.Query("SELECT * FROM scheduler ORDER BY date LIMIT :limit",
			sql.Named("limit", 15))
	}

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
