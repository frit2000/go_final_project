package main

import (
	// "time"
	"github.com/frit2000/go_final_project/db"
	"github.com/frit2000/go_final_project/httpServer"
)

func main() {

	db.DbExistance()
	httpServer.StartWebServer()
	// s := "5"
	// fmt.Printf("s=%02s", s)
	// dateInTimeFormat := time.Now().AddDate(1, 0, 0)
	// fmt.Println("after=", dateInTimeFormat)
}
