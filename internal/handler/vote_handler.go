package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"VK_API_BOT/internal/service"
)

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	userID := r.FormValue("user_id")

	args := strings.Split(text, " ")
	command := args[0]

	switch command {
	case "create":
		question := args[1]
		options := args[2:]

		id, err := service.CreatePoll(userID, question, options)
		if err != nil {
			http.Error(w, "failed to create poll", http.StatusInternalServerError)
			return
		}

		writeJSON(w, fmt.Sprintf("Poll created with ID: %s", id))
	default:
		writeJSON(w, "Unknown command")
	}
}

func writeJSON(w http.ResponseWriter, text string) {
	resp := map[string]string{
		"response_type": "in_channel",
		"text":          text,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
