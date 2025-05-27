package handlers

import (
	"encoding/json"
	"net/http"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
    messages := []string{"Пример сообщения 1", "Пример сообщения 2"}
    json.NewEncoder(w).Encode(messages)
}