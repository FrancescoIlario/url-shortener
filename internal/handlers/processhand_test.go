package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/FrancescoIlario/url-shortener/internal/handlers"
	"github.com/FrancescoIlario/url-shortener/internal/metrics"
	"github.com/FrancescoIlario/url-shortener/internal/mocks"
)

func Test_ProcessServeHTTP(t *testing.T) {
	repo := mocks.NewRepositoryGet(func(string) ([]byte, error) {
		return []byte("https://google.it/"), nil
	})
	metrics := &mocks.MetricsRepository{
		SaveAccessFunc:   func(url string, t time.Time) error { return nil },
		ReadAccessesFunc: func(id string) (*metrics.Accesses, error) { return nil, nil },
	}
	h := handlers.NewProcessHandler(repo, metrics)

	req, _ := http.NewRequest(http.MethodGet, "/c3b8f680a", nil)
	res := httptest.NewRecorder()

	h.ServeHTTP(res, req)

	if obt, exp := res.Result().StatusCode, http.StatusMovedPermanently; obt != exp {
		t.Errorf("expected status code %v, obtained %v", exp, obt)
	}
}
