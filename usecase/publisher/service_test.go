package publisher

import (
	"fmt"
	"testing"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/entity"

	"github.com/stretchr/testify/assert"
)

func newFixturePublisher() *entity.Publisher {
	return &entity.Publisher{
		Name:      "I Am Ozzy",
		Address:   "Ozzy Osbourne",
		CreatedAt: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u := newFixturePublisher()
	id, err := m.CreatePublisher(u.Name, u.Address)
	fmt.Println(id)
	assert.Nil(t, err)
	assert.False(t, u.CreatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u1 := newFixturePublisher()
	u2 := newFixturePublisher()
	u2.Name = "I Am Ozzy2"
	u1.Name = "I Am Ozzy"

	uID, _ := m.CreatePublisher(u1.Name, u1.Address)
	// uID2, _ := m.CreatePublisher(u2.Name, u2.Address)
	fmt.Println(uID)
	// fmt.Println(uID2)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchPublishers("ozzy")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "I Am Ozzy", c[0].Name)

		c, err = m.SearchPublishers("dio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := m.ListPublishers()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetPublisher(uID)
		fmt.Println(saved)
		fmt.Println(u1.Name)
		assert.Nil(t, err)
		assert.Equal(t, u1.Name, saved.Name)
	})
}

func Test_Update(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u := newFixturePublisher()
	id, err := m.CreatePublisher(u.Name, u.Address)
	assert.Nil(t, err)
	saved, _ := m.GetPublisher(id)
	saved.Name = "Lemmy: Biography"
	assert.Nil(t, m.UpdatePublisher(saved))
	updated, err := m.GetPublisher(id)
	assert.Nil(t, err)
	assert.Equal(t, "Lemmy: Biography", updated.Name)
}

func TestDelete(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u1 := newFixturePublisher()
	u2 := newFixturePublisher()
	u2ID, _ := m.CreatePublisher(u2.Name, u2.Address)

	err := m.DeletePublisher(u1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeletePublisher(u2ID)
	assert.Nil(t, err)
	_, err = m.GetPublisher(u2ID)
	assert.Equal(t, entity.ErrNotFound, err)
}
