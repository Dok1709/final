package main

import (
	"log"
	"net/http"
	"os"

	"todo-app/pkg/api"
	"todo-app/pkg/db"
)

func main() {
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	defer db.GetDB().Close()

	// Инициализация API эндпоинтов (Шаг 3)
	api.Init()

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	webDir := "./web"
	fileServer := http.FileServer(http.Dir(webDir))
	http.Handle("/", fileServer)

	log.Printf("Сервер запущен на порту %s", port)
	err = http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
}
