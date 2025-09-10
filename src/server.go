package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Заглушка, но вроде всё готово")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Вроде всё ок, мы рэди")
}

func main() {
	// Регистрируем обработчики для разных маршрутов
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/ready", readyHandler)
	http.HandleFunc("/helath", healthHandler)

	// Запускаем сервер
	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
