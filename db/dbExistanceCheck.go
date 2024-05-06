package db

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func DbExistance() {
	// appPath, err := os.Executable()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// envFile := os.Getenv("TODO_DBFILE")
	// if len(envFile) > 0 {
	// 	fmt.Println("зашли в env")
	// 	appPath = envFile
	// }

	//dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
	dbFile := "scheduler.db"
	_, err := os.Stat(dbFile)

	if err != nil {
		dbCreate()
		log.Println("Создана новая база данных с таблицей scheduler")
		return
	}
	log.Println("База данных уже существует")

}

func dbCreate() {
	db, err := sql.Open("sqlite", "scheduler.db")
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
