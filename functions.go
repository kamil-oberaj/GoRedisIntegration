package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (server *Server) PingOrPanic(ctx *context.Context) {
	if server.rdb == nil {
		panic("Redis client is nil")
	}

	pong, err := server.rdb.Ping(*ctx).Result()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Redis client is connected: %s\n", pong)
}

func NewRedisClient(url, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       db,
	})
}

func (server *Server) SetPerson(ctx *context.Context, person Person) {
	if server.rdb == nil {
		panic("Redis client is nil")
	}

	key := fmt.Sprintf("person:%s", person.ID)
	value, err := json.Marshal(person)

	if err != nil {
		panic(err)
	}

	_, err = server.rdb.Set(*ctx, key, value, 0).Result()

	if err != nil {
		panic(err)
	}
}

func (server *Server) GetAndPrintPerson(ctx *context.Context, id uuid.UUID) {
	if server.rdb == nil {
		panic("Redis client is nil")
	}

	key := fmt.Sprintf("person:%s", id)

	value, err := server.rdb.Get(*ctx, key).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Person: %s\n", value)
}
