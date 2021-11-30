package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/entity"
)

//PublisherMySQL mysql repo
type PublisherMySQL struct {
	db *sql.DB
}

//NewPublisherMySQL create new repository
func NewPublisherMySQL(db *sql.DB) *PublisherMySQL {
	return &PublisherMySQL{
		db: db,
	}
}

//Create a Publisher
func (r *PublisherMySQL) Create(e *entity.Publisher) (int64, error) {
	stmt, err := r.db.Prepare(`
		insert into publisher (name, address, created_at) 
		values(?,?,?)`)
	if err != nil {
		return e.ID, err
	}
	res, err := stmt.Exec(
		e.Name,
		e.Address,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	e.ID, _ = res.LastInsertId()
	return e.ID, nil
}

//Get a Publisher
func (r *PublisherMySQL) Get(id entity.ID) (*entity.Publisher, error) {
	stmt, err := r.db.Prepare(`select id, name, address, created_at from publisher where id = ?`)
	if err != nil {
		return nil, err
	}
	var b entity.Publisher
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&b.ID, &b.Name, &b.Address, &b.CreatedAt)
	}
	return &b, nil
}

//Update a Publisher
func (r *PublisherMySQL) Update(e *entity.Publisher) error {
	e.UpdatedAt = time.Now()
	fmt.Println(e)
	_, err := r.db.Exec("update publisher set name = ?, address = ?, updated_at = ? where id = ?", e.Name, e.Address, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

//Search Publishers
func (r *PublisherMySQL) Search(query string) ([]*entity.Publisher, error) {
	stmt, err := r.db.Prepare(`select id, name, address, created_at from publisher where name like ?`)
	if err != nil {
		return nil, err
	}
	var Publishers []*entity.Publisher
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b entity.Publisher
		err = rows.Scan(&b.ID, &b.Name, &b.Address, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		Publishers = append(Publishers, &b)
	}

	return Publishers, nil
}

//List Publishers
func (r *PublisherMySQL) List() ([]*entity.Publisher, error) {
	stmt, err := r.db.Prepare(`select id, name, address, created_at from publisher`)
	if err != nil {
		return nil, err
	}
	var Publishers []*entity.Publisher
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b entity.Publisher
		err = rows.Scan(&b.ID, &b.Name, &b.Address, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		Publishers = append(Publishers, &b)
	}
	return Publishers, nil
}

//Delete a Publisher
func (r *PublisherMySQL) Delete(id entity.ID) error {
	_, err := r.db.Exec("delete from publisher where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
