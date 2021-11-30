package publisher

import (
	"fmt"
	"strings"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/entity"
)

//Service Publisher usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreatePublisher create a Publisher
func (s *Service) CreatePublisher(name, address string) (int64, error) {
	b, err := entity.NewPublisher(name, address)
	if err != nil {
		return b.ID, err
	}
	return s.repo.Create(b)
}

//GetPublisher get a Publisher
func (s *Service) GetPublisher(id entity.ID) (*entity.Publisher, error) {
	b, err := s.repo.Get(id)
	fmt.Println(b)
	if b == nil || b.ID == 0 {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return b, nil
}

//SearchPublishers search Publishers
func (s *Service) SearchPublishers(query string) ([]*entity.Publisher, error) {
	Publishers, err := s.repo.Search(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(Publishers) == 0 {
		return nil, entity.ErrNotFound
	}
	return Publishers, nil
}

//ListPublishers list Publishers
func (s *Service) ListPublishers() ([]*entity.Publisher, error) {
	Publishers, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(Publishers) == 0 {
		return nil, entity.ErrNotFound
	}
	return Publishers, nil
}

//DeletePublisher Delete a Publisher
func (s *Service) DeletePublisher(id entity.ID) error {
	_, err := s.GetPublisher(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//UpdatePublisher Update a Publisher
func (s *Service) UpdatePublisher(e *entity.Publisher) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
