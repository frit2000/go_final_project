package server

import (
	"encoding/json"
	"net/http"
)

func (t TaskStore) delTask(w http.ResponseWriter, r *http.Request) {
	// var checkedID string
	var respTask RespTaskError

	id := r.FormValue("id")

	// err := t.db.QueryRow("SELECT id FROM scheduler WHERE id = :id", sql.Named("id", id)).Scan(&checkedID)
	// if err != nil {
	// 	log.Println("ошибка чтении данных по id:", err)
	// }

	// //проверяем, есть ли такой ID задачи
	// if len(checkedID) == 0 {
	// 	respTask.Err = "Ошибка, нет такого ID"
	// }

	// _, err = t.db.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
	// if err != nil {
	// 	log.Println("ошибка при обновлении записи БД:", err)
	// }

	resp, err := json.Marshal(&respTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}