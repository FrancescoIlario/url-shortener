package metrics

import (
	"database/sql"
	"time"

	// Postgres driver for sql driver
	_ "github.com/lib/pq"
)

type psqlRepository struct {
	connStr string
	db      *sql.DB
}

// NewPSQLRepository Postgres implementation for the metrics repository
func NewPSQLRepository(connStr string) (Repository, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	repo := psqlRepository{
		connStr: connStr,
		db:      db,
	}
	return &repo, nil
}

// SaveAccess save access to metrics db
func (r *psqlRepository) SaveAccess(urlID string, t time.Time) error {
	sqlInsert := `INSERT INTO Accesses(url_id, access_time) VALUES ($1, $2) RETURNING id`

	id := 0
	err := r.db.QueryRow(sqlInsert, urlID, t).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (r *psqlRepository) ReadAccesses(id string) (*Accesses, error) {
	rows, err := r.db.Query("SELECT * FROM GetAccesses($1);", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accesses Accesses
	if !rows.Next() {
		return nil, ErrIDNotFound
	}

	if err := rows.Scan(&accesses.OneDay, &accesses.OneWeek, &accesses.Total); err != nil {
		return nil, err
	}
	return &accesses, nil
}
