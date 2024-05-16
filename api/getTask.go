package api

import (
	"log"
	"net/http"
	"time"

	"github.com/frit2000/go_final_project/serverservice"
	"github.com/frit2000/go_final_project/servicetask"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	var tasks = map[string][]servicetask.Task{}
	var err error
	var s serverservice.ServerService

	searchString := r.FormValue("search")

	switch searchString {
	//если НЕ нажат поиск, то выбираем все записи
	case "":
		tasks, err = s.Server.GetAll()
		if err != nil {
			//http.Error(w, err.Error(), http.StatusBadRequest)
			log.Print("Ошибка валидации запроса:", err)
		}
		// rows, err = t.db.Query("SELECT * FROM scheduler ORDER BY date LIMIT :limit",
		// 	sql.Named("limit", limit))
	//если нажат поиск, то выбираем записи согласно строке поиска
	default:
		searchDate, errParse := time.Parse("02.01.2006", searchString)
		//если в поиске дата
		if errParse == nil {
			tasks, err = s.Server.GetSearchDate(searchDate)
			if err != nil {
				//http.Error(w, err.Error(), http.StatusBadRequest)
				log.Print("Ошибка валидации запроса:", err)
			}
			// rows, err = t.db.Query("SELECT * FROM scheduler WHERE date = :searchString LIMIT :limit",
			// 	sql.Named("searchString", searchDate.Format(params.DFormat)),
			// 	sql.Named("limit", limit))
			//если в поиске НЕ дата
		} else {
			tasks, err = s.Server.GetSearch(searchString)
			if err != nil {
				//http.Error(w, err.Error(), http.StatusBadRequest)
				log.Print("Ошибка валидации запроса:", err)
			}
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

	s.Response(tasks, w)
	// resp, err := json.Marshal(&tasks)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// w.Header().Set("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// if _, err = w.Write(resp); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

}
