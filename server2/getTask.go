package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

func (t TaskStore) getTask(w http.ResponseWriter, r *http.Request) {
	var tasks = map[string][]Task{
		"tasks": {},
	}
	var task Task
	var rows *sql.Rows
	var err error

	const limit = 15

	searchString := r.FormValue("search")

	switch searchString {
	//если НЕ нажат поиск, то выбираем все записи
	case "":
		// rows, err = t.db.Query("SELECT * FROM scheduler ORDER BY date LIMIT :limit",
		// 	sql.Named("limit", limit))
	//если нажат поиск, то выбираем записи согласно строке поиска
	default:
		searchDate, errParse := time.Parse("02.01.2006", searchString)
		//если в поиске дата
		if errParse == nil {
			// rows, err = t.db.Query("SELECT * FROM scheduler WHERE date = :searchString LIMIT :limit",
			// 	sql.Named("searchString", searchDate.Format(params.DFormat)),
			// 	sql.Named("limit", limit))
			//если в поиске НЕ дата
		} else {
			// rows, err = t.db.Query("SELECT * FROM scheduler WHERE title LIKE :searchString OR comment LIKE :searchString ORDER BY date LIMIT :limit",
			// 	sql.Named("searchString", "%"+searchString+"%"),
			// 	sql.Named("limit", limit))
		}
	}

	//если в процессе любого из поисков возникла ошибка, логируем
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	//log.Println("Ошибка запроса в базу:", err)
	// 	return
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	if err = rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	tasks["tasks"] = append(tasks["tasks"], task)
	// }

	// if err = rows.Err(); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	resp, err := json.Marshal(&tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
