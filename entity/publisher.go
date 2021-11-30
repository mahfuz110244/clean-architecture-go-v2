package entity

import (
	"time"
)

//Publisher data
type Publisher struct {
	ID        int64
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewPublisher create a new Publisher
func NewPublisher(name string, address string) (*Publisher, error) {
	b := &Publisher{
		Name:      name,
		Address:   address,
		CreatedAt: time.Now(),
	}
	err := b.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return b, nil
}

//Validate validate Publisher
func (b *Publisher) Validate() error {
	if b.Name == "" || b.Address == "" {
		return ErrInvalidEntity
	}
	return nil
}
