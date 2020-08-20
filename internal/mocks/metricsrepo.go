package mocks

import (
	"time"

	"github.com/FrancescoIlario/url-shortener/internal/metrics"
)

// MetricsRepository ...
type MetricsRepository struct {
	SaveAccessFunc   func(url string, t time.Time) error
	ReadAccessesFunc func(id string) (*metrics.Accesses, error)
}

// SaveAccess ...
func (r *MetricsRepository) SaveAccess(url string, t time.Time) error {
	if r.SaveAccessFunc == nil {
		return ErrNotImplemented
	}

	return r.SaveAccessFunc(url, t)
}

// ReadAccesses ...
func (r *MetricsRepository) ReadAccesses(id string) (*metrics.Accesses, error) {
	if r.ReadAccessesFunc == nil {
		return nil, ErrNotImplemented
	}

	return r.ReadAccessesFunc(id)
}
