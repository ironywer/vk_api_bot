package main

import (
	"log"
	"net/http"

	"VK_API_BOT/internal/handler"
)

func main() {
	http.HandleFunc("/vote", handler.VoteHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
