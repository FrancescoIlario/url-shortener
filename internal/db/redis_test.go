package db_test

import (
	"testing"

	"github.com/FrancescoIlario/url-shortener/internal/db"
	"github.com/FrancescoIlario/url-shortener/internal/idgen"
)

var dbconfig = db.Config{
	Host: "redis",
	Db:   13,
}

func Test_DbConn_Integration(t *testing.T) {
	if _, err := db.NewRepository(dbconfig); err != nil {
		t.Fatalf("error creating a new repository: %v", err)
	}
}

func Test_SaveAndGet_Integration(t *testing.T) {
	repo, err := db.NewRepository(dbconfig)
	if err != nil {
		t.Fatalf("error creating a new repository: %v", err)
	}

	id, err := idgen.NewID()
	if err != nil {
		t.Fatalf("error creating a new ID: %v", err)
	}

	url := "http://hello.com/hello"
	if err := repo.Save(id, url); err != nil {
		t.Fatalf("error saving a new url: %v", err)
	}

	if _, err := repo.Get(id); err != nil {
		t.Fatalf("error retrieving the url: %v", err)
	}
}
