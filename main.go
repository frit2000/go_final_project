package main

import (
	// "fmt"
	// "time"
	"github.com/frit2000/go_final_project/db"
	"github.com/frit2000/go_final_project/httpServer"
)

func main() {

	db.DbExistance()
	httpServer.StartWebServer()
	// fmt.Println("now=", time.Now())
	// dateInTimeFormat := time.Now().AddDate(1, 0, 0)
	// fmt.Println("after=", dateInTimeFormat)
}
