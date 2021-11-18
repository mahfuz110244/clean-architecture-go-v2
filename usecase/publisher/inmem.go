package publisher

import (
	"strings"

	"github.com/eminetto/clean-architecture-go-v2/entity"
)

//inmem in memory repo
type inmem struct {
	m map[entity.ID]*entity.Publisher
}

//newInmem create new repository
func newInmem() *inmem {
	var m = map[entity.ID]*entity.Publisher{}
	return &inmem{
		m: m,
	}
}

//Create a Publisher
func (r *inmem) Create(e *entity.Publisher) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get a Publisher
func (r *inmem) Get(id entity.ID) (*entity.Publisher, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update a Publisher
func (r *inmem) Update(e *entity.Publisher) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search Publishers
func (r *inmem) Search(query string) ([]*entity.Publisher, error) {
	var d []*entity.Publisher
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}
	return d, nil
}

//List Publishers
func (r *inmem) List() ([]*entity.Publisher, error) {
	var d []*entity.Publisher
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete a Publisher
func (r *inmem) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
