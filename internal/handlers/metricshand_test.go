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

func Test_MetricsServeHTTP(t *testing.T) {
	metrics := &mocks.MetricsRepository{
		SaveAccessFunc: func(url string, t time.Time) error { return nil },
		ReadAccessesFunc: func(id string) (*metrics.Accesses, error) {
			return &metrics.Accesses{OneDay: 1, OneWeek: 10, Total: 11}, nil
		},
	}
	h := handlers.NewMetricsHandler(metrics)

	req, _ := http.NewRequest(http.MethodGet, "/metrics/c3b8f680a", nil)
	res := httptest.NewRecorder()

	h.ServeHTTP(res, req)

	if obt, exp := res.Result().StatusCode, http.StatusOK; obt != exp {
		t.Errorf("expected status code %v, obtained %v", exp, obt)
	}
}
