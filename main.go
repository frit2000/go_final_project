package main

import (
	"github.com/frit2000/go_final_project/db"
	"github.com/frit2000/go_final_project/env"
	"github.com/frit2000/go_final_project/httpServer"
)

func main() {
	env.SetFlagParams()
	db.DbExistance()
	httpServer.StartWebServer()
}
