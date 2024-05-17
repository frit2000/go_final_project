package api

import (
	"net/http"
	"strconv"
)

func (srv Server) GetOneTask(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	task, tr, err := srv.Server.SrvService.GetOneTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if tr.Err != "" {
		srv.Server.Response(tr, w)
		return
	}

	srv.Server.Response(task, w)

}
