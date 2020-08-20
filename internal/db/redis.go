package db

import (
	"log"

	redis "github.com/alphazero/Go-Redis"
)

// Repository repository for shortened urls
type Repository interface {
	Get(id string) (string, error)
	Save(id string, url string) error
}

type repository struct {
	client redis.Client
}

// Config repository configuration structure
type Config struct {
	Host     string
	Db       int
	Password string
}

// NewRepository repository constructor
func NewRepository(c Config) (Repository, error) {
	spec := redis.DefaultSpec().Host(c.Host).Db(c.Db)
	client, e := redis.NewSynchClientWithSpec(spec)
	if e != nil {
		log.Println("failed to create the client", e)
		return nil, e
	}

	return &repository{
		client: client,
	}, nil
}

func (r *repository) Get(id string) (string, error) {
	value, e := r.client.Get(id)
	if e != nil {
		log.Println("error on Get", e)
		return "", e
	}
	return string(value), nil
}

func (r *repository) Save(id string, url string) error {
	return r.client.Set(id, []byte(url))
}
