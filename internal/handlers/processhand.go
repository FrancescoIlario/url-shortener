package handlers

import (
	"net/http"
	"time"

	"github.com/FrancescoIlario/url-shortener/internal/db"
	"github.com/FrancescoIlario/url-shortener/internal/metrics"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type processHandler struct {
	repo        db.Repository
	metricsRepo metrics.Repository
}

// NewProcessHandler process handler constructor
func NewProcessHandler(repo db.Repository, metricsRepo metrics.Repository) http.Handler {
	return &processHandler{
		repo:        repo,
		metricsRepo: metricsRepo,
	}
}

// ProcessURLHandler ...
func (h *processHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	logrus.Debugf("serving id %v", id)

	url, err := h.repo.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)

	go h.updateMetrics(id)
}

func (h *processHandler) updateMetrics(id string) {
	if err := h.metricsRepo.SaveAccess(id, time.Now()); err != nil {
		logrus.Errorf("error updating the metrics info: %v", err)
	}
}
