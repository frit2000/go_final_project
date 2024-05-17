package api

import (
	"net/http"
)

func (srv Server) DoneTask(w http.ResponseWriter, r *http.Request) {

	id := srv.Server.RequestId(r)

	task, err := srv.Server.SrvService.Done(id)
	checkErr(err)

	srv.Server.Response(task, w)

}
