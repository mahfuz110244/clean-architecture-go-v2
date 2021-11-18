package presenter

import (
	"github.com/eminetto/clean-architecture-go-v2/entity"
)

//Publisher data
type Publisher struct {
	ID      entity.ID `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
}
