package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/frit2000/go_final_project/api"
	"github.com/frit2000/go_final_project/env"
	"github.com/frit2000/go_final_project/serverservice"
	"github.com/frit2000/go_final_project/servicetask"

	"github.com/go-chi/chi"
)

func NewTaskStore(Db *sql.DB) servicetask.TaskStore {
	return servicetask.TaskStore{Db: Db}
}

func NewServerService(server servicetask.TaskStore) serverservice.ServerService {
	return serverservice.ServerService{Server: server}
}

func StartWebServer() {

	// подключаемся к БД
	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		log.Println("ошибка при подключении к БД:", err)
	}
	defer db.Close()

	store := NewTaskStore(db)
	server := NewServerService(store)

	todoPort := env.SetPort()
	webDir := "web"

	r := chi.NewRouter()
	r.Get("/api/nextdate", api.GetNextDate)
	r.Handle("/*", http.FileServer(http.Dir(webDir)))
	r.Post("/api/task", api.Auth(api.AddTask))
	r.Post("/api/task/done", api.Auth(api.DoneTask))
	r.Get("/api/task", api.Auth(api.GetOneTask))
	r.Delete("/api/task", api.Auth(api.DelTask))
	r.Put("/api/task", api.Auth(api.UpdTask))
	r.Get("/api/tasks", api.Auth(api.GetTask))
	r.Post("/api/signin", api.CheckPass)

	log.Println("Запускаем веб сервер")
	err = http.ListenAndServe(":"+todoPort, r)
	if err != nil {
		log.Println("ошибка при запуске веб сервера:", err)
	}
}
