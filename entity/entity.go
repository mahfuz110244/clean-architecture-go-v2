package entity

import (
	"math/rand"
)

// import "github.com/google/uuid"

//ID entity ID
type ID = int64

//NewID create a new entity ID
func NewID() ID {
	return ID(rand.Intn(1000000))
}

// //StringToID convert a string to an entity ID
// func StringToID(s string) (ID, error) {
// 	id, err := uuid.Parse(s)
// 	return ID(id), err
// }
