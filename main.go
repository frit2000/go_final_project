package main

import (
	"log"

	"github.com/frit2000/go_final_project/db"
	"github.com/frit2000/go_final_project/env"
	"github.com/frit2000/go_final_project/server"
)

func main() {
	env.SetFlagParams()

	err := db.DbExistance()
	if err != nil {
		log.Println("Ошибка при подключении к базе:", err)
		return
	}
	server.StartWebServer()
}
