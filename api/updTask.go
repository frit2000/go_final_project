package api

import (
	"net/http"
)

func (srv Server) UpdTask(w http.ResponseWriter, r *http.Request) {

	task, err := srv.Server.RequestUpd(r)
	checkErr(err)

	tr, err := srv.Server.ReqValidate(&task)
	checkErr(err)

	if tr.Err != "" {
		srv.Server.Response(tr, w)
		return
	}

	tr, err = srv.Server.SrvService.Update(task)
	checkErr(err)

	if tr.Err != "" {
		srv.Server.Response(tr, w)
		return
	}
	srv.Server.Response(task, w)
}
