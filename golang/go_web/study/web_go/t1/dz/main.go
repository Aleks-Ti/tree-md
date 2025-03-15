package main

import (
	"net/http"
	"strconv"
	"time"
)

var courses = map[int64]string{
	1: "Introduction to programming",
	2: "Introduction to algorithms",
	3: "Data structures",
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/courses/description", CourseDescHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Go to /courses/description"))
}

func CourseDescHandler(w http.ResponseWriter, r *http.Request) {
	// BEGIN (write your solution here)
	str_course_id := r.URL.Query().Get("course_id")

	course_id, err := strconv.ParseInt(str_course_id, 10, 64)
	if err != nil {
		http.Error(w, "Не допустимы запрос", http.StatusBadRequest)
		return
	}

	if desc_course, ok := courses[course_id]; ok {
		w.Write([]byte(desc_course))
	} else {
		http.Error(w, "Курс не найден", http.StatusNotFound)
	}
	// END
}
