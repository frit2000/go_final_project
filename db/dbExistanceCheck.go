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
	// envFile := os.Getenv("TODO_DBFILE")
	// if len(envFile) > 0 {
	// 	fmt.Println("зашли в env")
	// 	appPath = envFile
	// }

	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
	_, err = os.Stat(dbFile)

	if err != nil {
		dbCreate()
	}
}

func dbCreate() {
	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE scheduler (id INTEGER PRIMARY KEY AUTOINCREMENT,	
											date CHAR(8) NOT NULL DEFAULT "", 
											title VARCHAR(128) NOT NULL DEFAULT "", 
											comment VARCHAR(256) NOT NULL DEFAULT "",
											repeat VARCHAR(128))`)
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

}
