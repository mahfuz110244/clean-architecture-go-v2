package entity_test

import (
	"testing"

	"github.com/eminetto/clean-architecture-go-v2/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewPublisher(t *testing.T) {
	name := "Anupam"
	address := "Bangla Bazar, Dhaka"
	b, err := entity.NewPublisher(name, address)
	assert.Nil(t, err)
	assert.Equal(t, b.Name, name)
	assert.Equal(t, b.Address, address)
	assert.NotNil(t, b.ID)
}

func TestPublisherValidate(t *testing.T) {
	type test struct {
		name    string
		address string
		want    error
	}

	tests := []test{
		{
			name:    "Rokomari",
			address: "Dhaka",
			want:    nil,
		},
		{
			name:    "Rokomari",
			address: "",
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "",
			address: "Dhaka",
			want:    entity.ErrInvalidEntity,
		},
		{
			name:    "",
			address: "",
			want:    entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {

		_, err := entity.NewPublisher(tc.name, tc.address)
		assert.Equal(t, err, tc.want)
	}

}
