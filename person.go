package main

import (
	"github.com/google/uuid"
)

type Person struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}
