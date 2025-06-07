package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/choorka/backend/telegram"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func latestPostHandler(w http.ResponseWriter, r *http.Request) {
    data, err := telegram.GetLatestPostParsed()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/latest", latestPostHandler)

	handler := CORSMiddleware(mux)

	fmt.Println("Сервер запущен")
	http.ListenAndServe(":8080", handler)
}