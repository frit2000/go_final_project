package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/frit2000/go_final_project/env"
	"github.com/go-chi/chi"
)

func NewTaskStore(db *sql.DB) TaskStore {
	return TaskStore{db: db}
}

func StartWebServer() {

	// подключаемся к БД
	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		log.Println("ошибка при подключении к БД:", err)
	}
	defer db.Close()

	store := NewTaskStore(db)

	todoPort := env.SetPort()
	webDir := "web"

	r := chi.NewRouter()
	r.Get("/api/nextdate", getNextDate)
	r.Handle("/*", http.FileServer(http.Dir(webDir)))
	r.Post("/api/task", Auth(store.addTask))
	r.Post("/api/task/done", Auth(store.doneTask))
	r.Get("/api/task", Auth(store.getOneTask))
	r.Delete("/api/task", Auth(store.delTask))
	r.Put("/api/task", Auth(store.updTask))
	r.Get("/api/tasks", Auth(store.getTask))
	r.Post("/api/signin", checkPass)

	log.Println("Запускаем веб сервер")
	err = http.ListenAndServe(":"+todoPort, r)
	if err != nil {
		log.Println("ошибка при запуске веб сервера:", err)
	}
}
