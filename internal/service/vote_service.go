package service

import (
	"VK_API_BOT/internal/model"
	"VK_API_BOT/internal/storage"

	"github.com/google/uuid"
)

func CreatePoll(userID, question string, options []string) (string, error) {
	id := uuid.New().String()

	poll := model.Poll{
		ID:        id,
		CreatorID: userID,
		Question:  question,
		Options:   options,
		Votes:     make(map[string]string),
		IsClosed:  false,
	}

	if err := storage.SavePoll(poll); err != nil {
		return "", err
	}
	return id, nil
}
