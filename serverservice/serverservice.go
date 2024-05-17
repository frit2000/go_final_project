package serverservice

import (
	"encoding/json"
	"net/http"

	"github.com/frit2000/go_final_project/servicetask"
)

type ServerService struct {
	SrvService servicetask.TaskStore
}

func NewServerService(SrvService servicetask.TaskStore) ServerService {
	return ServerService{SrvService: SrvService}
}

// проверяем валидность запросов
func (ss ServerService) ReqValidate(t *servicetask.Task) (servicetask.TaskResp, error) {
	// проверяем что все поля date и title в task валидные
	var tr servicetask.TaskResp
	err := checkFieldsTask(t)
	if err != nil {
		tr.Err = "ошибка в формате поля date или title"
	}
	return tr, nil
}

func (ss ServerService) Response(t any, w http.ResponseWriter) {
	resp, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
