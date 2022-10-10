package model

import "github.com/google/uuid"

//NewID create a new entity ID
func NewID() string {
	return uuid.New().String()
}
