package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/frit2000/go_final_project/env"
	"github.com/frit2000/go_final_project/serverservice"
	"github.com/frit2000/go_final_project/servicetask"
	"github.com/go-chi/chi"
)

type Server struct {
	Server serverservice.ServerService
}

func NewServer(server serverservice.ServerService) Server {
	return Server{Server: server}
}

func StartWebServer() {

	// подключаемся к БД
	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		log.Println("ошибка при подключении к БД:", err)
	}
	defer db.Close()

	store := servicetask.NewTaskStore(db)
	service := serverservice.NewServerService(store)
	server := NewServer(service)

	todoPort := env.SetPort()
	webDir := "web"

	r := chi.NewRouter()
	r.Get("/api/nextdate", GetNextDate)
	r.Handle("/*", http.FileServer(http.Dir(webDir)))
	r.Post("/api/task", (server.AddTask))
	r.Post("/api/task/done", (server.DoneTask))
	r.Get("/api/task", (server.GetOneTask))
	r.Delete("/api/task", (server.DelTask))
	r.Put("/api/task", (server.UpdTask))
	r.Get("/api/tasks", (server.GetTask))
	//	r.Post("/api/signin", CheckPass)

	log.Println("Запускаем веб сервер")
	err = http.ListenAndServe(":"+todoPort, r)
	if err != nil {
		log.Println("ошибка при запуске веб сервера:", err)
	}
}
