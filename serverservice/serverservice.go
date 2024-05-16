package serverservice

import (
	"encoding/json"
	"net/http"

	"github.com/frit2000/go_final_project/servicetask"
)

type ServerService struct {
	Server servicetask.TaskStore
}

func NewServerService(server servicetask.TaskStore) ServerService {
	return ServerService{Server: server}
}

// проверяем валидность запросов
func (ss ServerService) ReqValidate(t *servicetask.Task) error {
	// проверяем что все поля date и title в task валидные
	var task = servicetask.Task{}
	err := checkFieldsTask(t)
	if err != nil {
		task.Err = "ошибка в формате поля date или title"
	}
	return nil
}

func (ss ServerService) Response(t any, w http.ResponseWriter) {
	resp, err := json.Marshal(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if _, err = w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		// log.Println("Не удалось записать данные в html:", err)
		// return
	}
}
