package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/choorka/backend/telegram"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
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
    http.HandleFunc("/latest", latestPostHandler)
    fmt.Println("Сервер запущен")
    http.ListenAndServe(":8080", CORSMiddleware())
}
