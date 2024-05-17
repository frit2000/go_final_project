package api

import (
	"net/http"
)

func (srv Server) DelTask(w http.ResponseWriter, r *http.Request) {

	id := srv.Server.RequestId(r)

	tr, err := srv.Server.SrvService.Delete(id)
	checkErr(err)

	srv.Server.Response(tr, w)

}
