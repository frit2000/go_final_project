package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/frit2000/go_final_project/servicetask"
)

func (srv Server) UpdTask(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var task servicetask.Task
	var tr servicetask.TaskResp

	// получаем данные из веб-интерфейса
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//переводим данные в стркутуру task
	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		log.Println(err)
	}

	tr, err = srv.Server.ReqValidate(&task)
	if err != nil {
		log.Println(err)
	}

	if tr.Err != "" {
		srv.Server.Response(tr, w)
		return
	}

	tr, err = srv.Server.SrvService.Update(task)
	if err != nil {
		log.Println(err)
	}

	if tr.Err != "" {
		srv.Server.Response(tr, w)
		return
	}
	srv.Server.Response(task, w)
}
