package metrics

import (
	"fmt"
	"time"
)

// Repository interface for metrics repository
type Repository interface {
	SaveAccess(id string, t time.Time) error
	ReadAccesses(id string) (*Accesses, error)
}

// Accesses ...
type Accesses struct {
	OneDay  int `json:"OneDay"`
	OneWeek int `json:"OneWeek"`
	Total   int `json:"Total"`
}

// ErrIDNotFound error returned when the id is not found
var ErrIDNotFound = fmt.Errorf("Id not found")
