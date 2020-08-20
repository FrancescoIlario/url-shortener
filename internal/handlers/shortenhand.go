package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/FrancescoIlario/url-shortener/internal/db"
	"github.com/FrancescoIlario/url-shortener/internal/idgen"
	"github.com/sirupsen/logrus"
)

type shortenHandler struct {
	address string
	repo    db.Repository
}

// ShortenURLRequest the Shorten request
type ShortenURLRequest struct {
	URL string `json:"Url"`
}

// ShortenURLResponse the payload in the response to the Shorten request
type ShortenURLResponse struct {
	ID  string `json:"Id"`
	URL string `json:"Url"`
}

// NewShortenHandler shorten handler constructor
func NewShortenHandler(address string, repo db.Repository) http.Handler {
	return &shortenHandler{
		address: address,
		repo:    repo,
	}
}

func (h *shortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := h.parseRequest(r)
	if err != nil {
		logrus.Infof("error parsing the request: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// check input validity
	if !isValidURL(req.URL) {
		logrus.Infof("invalid url provided: %v", req.URL)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Produce a new Id
	id, err := idgen.NewID()
	if err != nil {
		logrus.Errorf("error generating a new string ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Save to db
	h.repo.Save(id, req.URL)

	// Produce response
	url := h.composeAddress(&id)
	logrus.Debugf("build an URL: \"%v\"", url)
	response := ShortenURLResponse{
		ID:  id,
		URL: url,
	}

	payload, err := json.Marshal(response)
	if err != nil {
		logrus.Errorf("error marshaling response to JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(payload)
}

func (h *shortenHandler) parseRequest(r *http.Request) (*ShortenURLRequest, error) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var req ShortenURLRequest
	if err := json.Unmarshal(reqData, &req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (h *shortenHandler) composeAddress(id *string) string {
	return h.address + *id
}

func isValidURL(urlToTest string) bool {
	if _, err := url.ParseRequestURI(urlToTest); err != nil {
		return false
	}

	if u, err := url.Parse(urlToTest); err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
