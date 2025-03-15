package main

import (
	"net/http"

	logrus "github.com/sirupsen/logrus"
)

func main() {
	// log.SetFormatter(&log.TextFormatter{
	// 	DisableColors: false,
	// 	FullTimestamp: true,
	// })

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Welcome to Hexlet"))
		if err != nil {
			// Ошибка логируется функцией WithError
			logrus.WithError(err).Error("write hello hexlet")
		}
	})

	port := "80"
	// Дополнительная информация передается функцией WithFields
	logrus.WithFields(logrus.Fields{
		"port": port,
	}).Info("Starting a web-server on port")
	logrus.Fatal(http.ListenAndServe(":"+port, nil))
}
