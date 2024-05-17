package api

import (
	"net/http"
	"time"

	"github.com/frit2000/go_final_project/servicetask"
)

func (srv Server) GetTask(w http.ResponseWriter, r *http.Request) {
	var tasks = map[string][]servicetask.Task{}
	var err error
	var tr servicetask.TaskResp

	searchString := r.FormValue("search")

	switch searchString {
	//если НЕ нажат поиск, то выбираем все записи
	case "":
		tasks, tr, err = srv.Server.SrvService.GetAll()
		checkErr(err)

	//если нажат поиск, то выбираем записи согласно строке поиска
	default:
		searchDate, errParse := time.Parse("02.01.2006", searchString)
		//если в поиске дата
		if errParse == nil {
			tasks, err = srv.Server.SrvService.GetSearchDate(searchDate)
			checkErr(err)

			//если в поиске НЕ дата
		} else {
			tasks, err = srv.Server.SrvService.GetSearch(searchString)
			checkErr(err)
		}
	}

	if tr.Err != "" {
		srv.Server.Response(tr, w)
		return
	}

	srv.Server.Response(tasks, w)

}
