package publisher

import (
	"github.com/eminetto/clean-architecture-go-v2/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Publisher, error)
	Search(query string) ([]*entity.Publisher, error)
	List() ([]*entity.Publisher, error)
}

//Writer Publisher writer
type Writer interface {
	Create(e *entity.Publisher) (entity.ID, error)
	Update(e *entity.Publisher) error
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetPublisher(id entity.ID) (*entity.Publisher, error)
	SearchPublishers(query string) ([]*entity.Publisher, error)
	ListPublishers() ([]*entity.Publisher, error)
	CreatePublisher(name string, address string) (entity.ID, error)
	UpdatePublisher(e *entity.Publisher) error
	DeletePublisher(id entity.ID) error
}
