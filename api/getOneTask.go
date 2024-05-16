package api

import (
	"log"
	"net/http"

	"github.com/frit2000/go_final_project/serverservice"
	"github.com/frit2000/go_final_project/servicetask"
)

func GetOneTask(w http.ResponseWriter, r *http.Request) {
	var task servicetask.Task
	var s serverservice.ServerService

	task.Id = r.FormValue("id")

	if err := s.Server.Update(&task); err != nil {
		//ttp.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Ошибка валидации запроса:", err)
	}
	// err := t.db.QueryRow("SELECT * FROM scheduler WHERE id = :id", sql.Named("id", id)).Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	// if err != nil {
	// 	log.Println("ошибка чтении данных по id:", err)
	// }

	// //проверяем, есть ли такой ID задачи
	// if len(task.Id) == 0 {
	// 	respTask.Err = "Ошибка, нет такого ID"
	// }

	s.Response(task, w)
	// if respTask.Err == "" {
	// 	resp, err = json.Marshal(&task)
	// } else {
	// 	resp, err = json.Marshal(&respTask)
	// }
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
