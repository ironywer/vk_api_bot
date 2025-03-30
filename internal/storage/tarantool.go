package storage

import (
	"log"
	"os"

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

func SavePoll(p model.Poll) error {
	_, err := conn.Insert("polls", []interface{}{p.ID, p.CreatorID, p.Question, p.Options, p.Votes, p.IsClosed})
	return err
}
