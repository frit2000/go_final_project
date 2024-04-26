package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func DbExistance() {
	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	envFile := os.Getenv("TODO_DBFILE")
	fmt.Println("env=", envFile)
	if len(envFile) > 0 {
		appPath = envFile
	}

	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
	_, err = os.Stat(dbFile)

	//	var install bool
	if err != nil {
		dbCreate()
	}
	// если install равен true, после открытия БД требуется выполнить
	// sql-запрос с CREATE TABLE и CREATE INDEX
}

func dbCreate() {
	db, err := sql.Open("sqlite", "db/scheduler.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE scheduler (id INTEGER PRIMARY KEY AUTOINCREMENT,date CHAR(8),title VARCHAR(128),comment VARCHAR(256),repeat VARCHAR(128))")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec("CREATE INDEX dateindex ON scheduler (date)")
	if err != nil {
		fmt.Println(err)
		return
	}

	row := db.QueryRow("SELECT * FROM scheduler")
	fmt.Println("base result=", row)

}
