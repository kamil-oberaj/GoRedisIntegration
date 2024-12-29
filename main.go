package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Server struct {
	rdb *redis.Client
}

func main() {
	server := Server{
		rdb: NewRedisClient("localhost:6379", "", 0),
	}

	server.PingOrPanic(&ctx)

	id, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	person := Person{
		ID:   id,
		Name: "John Doe",
		Age:  30,
	}

	server.SetPerson(&ctx, person)
	server.GetAndPrintPerson(&ctx, id)
}
