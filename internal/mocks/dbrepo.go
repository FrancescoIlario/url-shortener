package mocks

import (
	"github.com/FrancescoIlario/url-shortener/internal/db"
)

type repository struct {
	get func(string) ([]byte, error)
	set func(key string, arg1 []byte) error
}

// NewRepositoryGetSet ...
func NewRepositoryGetSet(get func(string) ([]byte, error), set func(key string, arg1 []byte) error) db.Repository {
	return &repository{get: get, set: set}
}

// NewRepositoryGet ...
func NewRepositoryGet(get func(string) ([]byte, error)) db.Repository {
	return &repository{get: get}
}

// NewRepositorySet ...
func NewRepositorySet(set func(key string, arg1 []byte) error) db.Repository {
	return &repository{set: set}
}

func (r *repository) Get(id string) (string, error) {
	if r.get == nil {
		return "", ErrNotImplemented
	}

	value, e := r.get(id)
	if e != nil {
		return "", e
	}
	return string(value), nil
}

func (r *repository) Save(id string, url string) error {
	if r.set == nil {
		return ErrNotImplemented
	}

	if err := r.set(id, []byte(url)); err != nil {
		return err
	}

	return nil
}
