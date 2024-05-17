package api

import (
	"net/http"
)

func (srv Server) GetOneTask(w http.ResponseWriter, r *http.Request) {

	id := srv.Server.RequestId(r)

	task, tr, err := srv.Server.SrvService.GetOneTask(id)
	checkErr(err)

	if tr.Err != "" {
		srv.Server.Response(tr, w)
		return
	}

	srv.Server.Response(task, w)

}
