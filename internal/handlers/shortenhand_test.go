package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FrancescoIlario/url-shortener/internal/handlers"
	"github.com/FrancescoIlario/url-shortener/internal/mocks"
)

func Test_ShortenServeHTTP_HappyPath(t *testing.T) {
	repo := mocks.NewRepositorySet(func(key string, arg1 []byte) error { return nil })
	h := handlers.NewShortenHandler("localhost:8080", repo)

	reqData := handlers.ShortenURLRequest{
		URL: "https://google.it",
	}
	b, err := json.Marshal(&reqData)
	if err != nil {
		t.Fatalf("error marshaling JSON: %v", err)
	}
	reader := bytes.NewReader(b)

	req, _ := http.NewRequest(http.MethodPost, "/shorten/anon", reader)
	res := httptest.NewRecorder()

	h.ServeHTTP(res, req)

	if obt, exp := res.Result().StatusCode, http.StatusOK; obt != exp {
		t.Errorf("expected status code %v, obtained %v", exp, obt)
	}
}

func Test_ShortenServeHTTP_BadURL_TableTest(t *testing.T) {
	repo := mocks.NewRepositorySet(func(key string, arg1 []byte) error { return nil })
	h := handlers.NewShortenHandler("localhost:8080", repo)

	urls := []string{
		"://ssa.it",
		"",
		"~/home",
		"index.html",
		"./here",
	}

	for _, u := range urls {
		reqData := handlers.ShortenURLRequest{
			URL: u,
		}
		b, err := json.Marshal(&reqData)
		if err != nil {
			t.Fatalf("error marshaling JSON: %v", err)
		}
		reader := bytes.NewReader(b)

		req, _ := http.NewRequest(http.MethodPost, "/shorten/anon", reader)
		res := httptest.NewRecorder()

		h.ServeHTTP(res, req)

		if obt, exp := res.Result().StatusCode, http.StatusBadRequest; obt != exp {
			t.Errorf("expected status code %v, obtained %v", exp, obt)
		}
	}
}
