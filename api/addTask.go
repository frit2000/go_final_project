package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/frit2000/go_final_project/serverservice"
	"github.com/frit2000/go_final_project/servicetask"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var task servicetask.Task
	var s serverservice.ServerService

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

	if err = s.ReqValidate(&task); err != nil {
		//ttp.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Ошибка валидации запроса:", err)
	}
	// // проверяем что все поля date и title в task валидные
	// err = checkFieldsTask(&task)
	// if err != nil {
	// 	respTaskAdd.Err = "ошибка в формате поля date или title"
	// }

	err = s.Server.Add(&task)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Ошибка валидации запроса:", err)
	}

	// //записываем поля структуры task в БД
	// res, err := t.db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
	// 	sql.Named("date", task.Date),
	// 	sql.Named("title", task.Title),
	// 	sql.Named("comment", task.Comment),
	// 	sql.Named("repeat", task.Repeat))
	// if err != nil {
	// 	log.Println("ошибка при добавлении записи БД:", err)
	// }

	// //получаем ID последней добавленной записи
	// lastID, err := res.LastInsertId()
	// if err != nil {
	// 	log.Println("ошибка получении последнего ID:", err)

	// }
	// respTaskAdd.Id = strconv.Itoa(int(lastID))

	s.Response(task, w)
	// resp, err := json.Marshal(&respTaskAdd)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	// if _, err = w.Write(resp); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
}