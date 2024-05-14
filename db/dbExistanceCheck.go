package db

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"

	"github.com/frit2000/go_final_project/env"
)

func DbExistance() {
	// dbFile := "scheduler.db"

	// envFile := os.Getenv("TODO_DBFILE")
	// if len(envFile) > 0 {
	// 	dbFile = filepath.Join(envFile, "scheduler.db")
	// }

	// log.Println("путь к БД:", dbFile)

	dbFile := env.DbName()
	_, err := os.Stat(dbFile)
	if err != nil {
		log.Println("Создаем новую базу данных с таблицей scheduler")
		dbCreate(dbFile)
		return
	}
	log.Println("База данных уже существует")
}

func dbCreate(dbFile string) {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		log.Println("ошибка при подключении к БД:", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE scheduler (id INTEGER PRIMARY KEY AUTOINCREMENT,	
											date CHAR(8), 
											title VARCHAR(128) NOT NULL DEFAULT "", 
											comment VARCHAR(256) ,
											repeat VARCHAR(128))`)
	if err != nil {
		log.Println("ошибка при создании таблицы в БД:", err)
	}

	_, err = db.Exec("CREATE INDEX dateindex ON scheduler (date)")
	if err != nil {
		log.Println("ошибка при создании индекса dateindex в БД:", err)
	}
}
