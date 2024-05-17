package api

import (
	"log"
	"net/http"
	"strconv"
)

func (srv Server) GetOneTask(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.FormValue("id"))

	task, tr, err := srv.Server.SrvService.GetOneTask(id)
	if err != nil {
		log.Println(err)
	}

	if tr.Err != "" {
		srv.Server.Response(tr, w)
		return
	}

	srv.Server.Response(task, w)

}
