package api

import (
	"fmt"
	"net/http"
	"strconv"
)

func (srv Server) DelTask(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	fmt.Println("id from html=", id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	tr, err := srv.Server.SrvService.Delete(idInt)
	if err != nil {
		fmt.Println("ошибка при удалении:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)

	}
	srv.Server.Response(tr, w)

}
