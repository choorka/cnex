package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/choorka/backend/telegram"
)

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
    http.HandleFunc("/latest", latestPostHandler)
    fmt.Println("Сервер запущен на http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
