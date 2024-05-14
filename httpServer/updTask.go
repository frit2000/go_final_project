package httpServer

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func (t TaskStore) updTask(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var task Task
	var respTaskAdd RespTaskError
	var count int

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
		respTaskAdd.Err = "ошибка в формате поля date или title"
	}

	//проверяем, есть ли такой ID задачи
	err = t.db.QueryRow("SELECT COUNT (*) FROM scheduler WHERE id = :id", sql.Named("id", task.Id)).Scan(&count)
	if err != nil {
		log.Println("ошибка чтении данных из БД:", err)
	}
	if count == 0 {
		respTaskAdd.Err = "задача не найдена"
	}
	//обновляем поля структуры task в БД
	_, err = t.db.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.Id))
	if err != nil {
		log.Println("ошибка при обновлении записи БД:", err)
	}

	//	rowsAffected, err := result.RowsAffected()

	resp, err := json.Marshal(&respTaskAdd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(resp); err != nil {
		log.Println("Не удалось записать данные в html:", err)
		return
	}
}
