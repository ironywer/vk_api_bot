package main

import (
	"VK_API_BOT/internal/handler"
	"VK_API_BOT/internal/middleware"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/vote", handler.VoteHandler)

	logged := middleware.Logger(mux)

	http.ListenAndServe(":8080", logged)
}
