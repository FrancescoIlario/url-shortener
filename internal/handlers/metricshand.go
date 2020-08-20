package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/FrancescoIlario/url-shortener/internal/metrics"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type metricsHandler struct {
	repo metrics.Repository
}

// NewMetricsHandler metrics handler constructor
func NewMetricsHandler(repo metrics.Repository) http.Handler {
	return &metricsHandler{
		repo: repo,
	}
}

func (h *metricsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	logrus.Debugf("serving metrics for id: %v", id)

	accesses, err := h.repo.ReadAccesses(id)
	if err != nil {
		logrus.Errorf("error reading accesses metrics for id %v: %v", id, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(accesses)
}
