package server

import (
	"github.com/frit2000/go_final_project/params"
	"github.com/frit2000/go_final_project/servicetask"
)

type ServerService struct {
	server servicetask.TaskStore
}

func NewServerService(server servicetask.TaskStore) ServerService {
	return ServerService{server: server}
}

// проверяем валидность запросов
func (ss ServerService) ReqValidate(t servicetask.Task) (params.RespTaskError, error) {
	// проверяем что все поля date и title в task валидные
	err := checkFieldsTask(&servicetask.t)
	if err != nil {
		respTaskAdd.Err = "ошибка в формате поля date или title"
	}
}

func (ss ServerService) Response(t servicetask.Task) (string, error) {

}
