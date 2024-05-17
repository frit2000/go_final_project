package api

import (
	"log"
	"net/http"
	"strconv"
)

func (srv Server) DelTask(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	idInt, _ := strconv.Atoi(id)

	tr, err := srv.Server.SrvService.Delete(idInt)
	if err != nil {
		log.Println(err)
	}

	srv.Server.Response(tr, w)

}
