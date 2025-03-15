package main

import (
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	cwd, _ := os.Getwd()
	logFile := filepath.Join(cwd, ".log")
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logger.Fatal(err)
	}
	defer file.Close()
	logger.SetOutput(file)

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Go to /sum"))
	})

	mux.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		// BEGIN (write your solution here)
		x := r.URL.Query().Get("x")
		int_x, err := strconv.Atoi(x)
		if err != nil {
			http.Error(w, "Не допустимый параметр X в запросе", http.StatusBadRequest)
			return
		}
		y := r.URL.Query().Get("y")
		int_y, err := strconv.Atoi(y)
		if err != nil {
			http.Error(w, "Не допустимый параметр Y в запросе", http.StatusBadRequest)
			return
		}
		var result int = int_x + int_y
		if (int_y > 0 && int_x > math.MaxInt-int_y) || (int_y < 0 && int_x < math.MinInt-int_y) {
			w.Write([]byte("-1"))
			logger.WithFields(logrus.Fields{
				"x": x,
				"y": y,
			}).Warn("Sum overflows int")
			return
		}
		resultStr := strconv.Itoa(result)
		w.Write([]byte((resultStr)))
		// END
	})
	port := "8080"
	logWithPort := logrus.WithFields(logrus.Fields{
		"port": port,
	})
	logWithPort.Info("Starting a web-server on port")
	logWithPort.Fatal(server.ListenAndServe())
}
