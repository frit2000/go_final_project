package httpServer

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/frit2000/go_final_project/nextdate"
	"github.com/go-chi/chi"
)

func getNextDate(w http.ResponseWriter, r *http.Request) {
	//получаем параметры из запроса и переводим now в формат времени
	nowInString := r.FormValue("now")
	now, err := time.Parse("20060102", nowInString)
	if err != nil {
		log.Println("ошибка парсинга формата заданной даты:", err)
	}
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	// получаем новую дату
	s, err := nextdate.NextDate(now, date, repeat)
	if err != nil {
		log.Println("ошибка при переносе даты:", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func StartWebServer() {

	todoPort := os.Getenv("TODO_PORT")
	if todoPort == "" {
		todoPort = "7540"
	}
	webDir := "web"

	r := chi.NewRouter()
	r.Get("/api/nextdate", getNextDate)
	fmt.Println("path=", http.FileServer(http.Dir(webDir)))
	//	r.Handle("/", http.FileServer(http.Dir(webDir)))
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	log.Println("Запускаем веб сервер")
	err := http.ListenAndServe(":"+todoPort, r)
	if err != nil {
		log.Println("ошибка при запуске веб сервера:", err)
	}
	log.Println("Веб сервер запущен")
}
