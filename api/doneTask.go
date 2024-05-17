package api

import (
	"net/http"
	"strconv"
)

func (srv Server) DoneTask(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	task, err := srv.Server.SrvService.Done(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	srv.Server.Response(task, w)

}
