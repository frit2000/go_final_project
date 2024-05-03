package httpServer

import (
	"log"
	"net/http"
	"time"

	"github.com/frit2000/go_final_project/nextdate"
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