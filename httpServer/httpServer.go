package httpServer

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func StartWebServer() {

	todoPort := os.Getenv("TODO_PORT")
	if todoPort == "" {
		todoPort = "7540"
	}
	webDir := "web"

	r := chi.NewRouter()
	r.Get("/api/nextdate", getNextDate)
	r.Handle("/*", http.FileServer(http.Dir(webDir)))
	r.Post("/api/task", addTask)

	log.Println("Запускаем веб сервер")
	err := http.ListenAndServe(":"+todoPort, r)
	if err != nil {
		log.Println("ошибка при запуске веб сервера:", err)
	}
	log.Println("Веб сервер запущен")
}
