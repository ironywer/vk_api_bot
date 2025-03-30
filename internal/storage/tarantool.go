package storage

import (
	"log"
	"os"
	"sync"
	"errors"
	"VK_API_BOT/internal/model"

	"github.com/tarantool/go-tarantool"
)

var conn *tarantool.Connection

func init() {
	var err error
	host := os.Getenv("TARANTOOL_HOST")
	if host == "" {
		host = "localhost"
	}
	conn, err = tarantool.Connect(host+":3301", tarantool.Opts{})
	if err != nil {
		log.Fatalf("Failed to connect to Tarantool: %s", err)
	}
}

var (
	pollStore  = make(map[string]model.Poll)
	storeMutex sync.Mutex
)

func SavePoll(poll model.Poll) error {
	storeMutex.Lock()
	defer storeMutex.Unlock()

	pollStore[poll.ID] = poll
	return nil
}

func GetPoll(id string) (model.Poll, error) {
	storeMutex.Lock()
	defer storeMutex.Unlock()

	poll, ok := pollStore[id]
	if !ok {
		return model.Poll{}, errors.New("poll not found")
	}
	return poll, nil
}

func UpdatePoll(poll model.Poll) error {
	storeMutex.Lock()
	defer storeMutex.Unlock()

	pollStore[poll.ID] = poll
	return nil
}

func DeletePoll(pollID string) error {
	storeMutex.Lock()
	defer storeMutex.Unlock()

	if _, exists := pollStore[pollID]; !exists {
		return errors.New("poll not found")
	}

	delete(pollStore, pollID)
	return nil
}