package api

import (
	"log"
	"net/http"
	"strconv"
)

func (srv Server) DoneTask(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println(err)
	}

	task, err := srv.Server.SrvService.Done(id)
	if err != nil {
		log.Println(err)
	}

	srv.Server.Response(task, w)

}
