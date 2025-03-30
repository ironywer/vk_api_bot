package service

import (
	"errors"
	"strings"

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

func CastVote(userID, pollID, option string) error {
	poll, err := storage.GetPoll(pollID)
	if err != nil {
		return err
	}

	if poll.IsClosed {
		return errors.New("poll is closed")
	}

	if _, voted := poll.Votes[userID]; voted {
		return errors.New("user has already voted")
	}

	valid := false
	for _, opt := range poll.Options {
		if strings.EqualFold(opt, option) {
			valid = true
			break
		}
	}
	if !valid {
		return errors.New("invalid option")
	}

	poll.Votes[userID] = option
	return storage.UpdatePoll(poll)
}