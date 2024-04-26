package httpServer

import (
	"fmt"
	"net/http"
	"os"
)

func StartWebServer() {
	todoPort := os.Getenv("TODO_PORT")
	if todoPort == "" {
		todoPort = "7540"
	}
	webDir := "web"
	http.Handle(`/`, http.FileServer(http.Dir(webDir)))
	err := http.ListenAndServe(":"+todoPort, nil)
	if err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
