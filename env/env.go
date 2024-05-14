package env

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

// определяем флаги
func SetFlagParams() {
	pass := flag.String("password", "", "Пароль для приложения")
	port := flag.String("port", "7540", "Порт для запуска веб сервера")
	dbPath := flag.String("dbpath", "", "Путь к базе данных")

	flag.Parse()
	os.Setenv("TODO_PASSWORD", *pass)
	os.Setenv("TODO_PORT", *port)
	os.Setenv("TODO_DBFILE", *dbPath)
}

func DbName() string {
	dbFile := "scheduler.db"
	envFile := os.Getenv("TODO_DBFILE")
	if len(envFile) > 0 {
		dbFile = filepath.Join(envFile, "scheduler.db")
	}
	log.Println("путь к БД:", dbFile)
	return dbFile
}

func SetPass() string {
	return os.Getenv("TODO_PASSWORD")
}

func SetPort() string {
	return os.Getenv("TODO_PORT")
}
