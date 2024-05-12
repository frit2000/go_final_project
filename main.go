package main

import (
	"os"

	"github.com/frit2000/go_final_project/db"
	"github.com/frit2000/go_final_project/httpServer"
)

func main() {
	os.Setenv("TODO_PASSOWRD", "123")
	db.DbExistance()
	httpServer.StartWebServer()
}
