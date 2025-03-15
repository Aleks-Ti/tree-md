package main

import (
	"fmt"
	"net/http"
)

func main() {
	// обозначаем, что на запрос по пути "/" возвращается строка "Hello World"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// тело ответа — это массив байт
		w.Write([]byte("Hello world!"))
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hexlet is the leading educational platform for Software Engineers"))
	})

	http.HandleFunc("/courses", func(response_writer http.ResponseWriter, request *http.Request) {
		// считываем параметр page из запроса
		page := request.URL.Query().Get("page")

		// рассчитываем, какую страницу нужно вернуть
		var pageCourses string
		switch page {
		case "":
			pageCourses = "Введите номер курса!"
		case "1":
			pageCourses = "Как написать свой первый \"Hello world\" на go..."
		case "2":
			pageCourses = "Как работает сборщик мусора в go..."
		default:
			pageCourses = "Курс в разработке..."
		}

		// возвращаем страницу курса
		response_writer.Write([]byte(pageCourses))
	})

	// запускаем веб-приложение для обработки запросов
	fmt.Println("Server start localhost:80")
	http.ListenAndServe(":80", nil)
}
