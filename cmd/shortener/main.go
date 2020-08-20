package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/FrancescoIlario/url-shortener/internal/db"
	"github.com/FrancescoIlario/url-shortener/internal/handlers"
	"github.com/FrancescoIlario/url-shortener/internal/metrics"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const addressEnvKey = "ADDRESS"
const psqlConnStrEnvKey = "POSTRGRES_CONNSTR"

const defaultAddress = "http://localhost:8080/"
const defaultPSQLConnStr = "host=postgres port=5432 user=postgres password=supersecret dbname=metrics sslmode=disable"

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableLevelTruncation: true,
		DisableTimestamp:       false,
		ForceColors:            true,
		FullTimestamp:          true,
		DisableColors:          false,
	})

	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	logrus.Debugln("starting url shortener")

	// parsing configuration
	address := getAddressFromEnv()
	psqlConnStr := getPSQLConnStrFromEnv()

	// set up repositories
	repo, err := db.NewRepository(db.Config{
		Host: "redis",
		Db:   13,
	})
	if err != nil {
		return fmt.Errorf("can not connect to redis db: %w", err)
	}
	metricsRepo, err := metrics.NewPSQLRepository(psqlConnStr)
	if err != nil {
		return fmt.Errorf("can not connect to metrics db: %w", err)
	}

	// setup router
	r := mux.NewRouter()
	r.Handle("/{id}", handlers.NewProcessHandler(repo, metricsRepo)).Methods(http.MethodGet)
	r.Handle("/shorten/anon", handlers.NewShortenHandler(address, repo)).Methods(http.MethodPost)
	r.Handle("/metrics/{id}", handlers.NewMetricsHandler(metricsRepo)).Methods(http.MethodGet)
	http.Handle("/", r)

	// start serving
	if err := http.ListenAndServe(":8080", r); err != nil {
		return fmt.Errorf("error serving: %w", err)
	}
	return nil
}

func getAddressFromEnv() string {
	addr := os.Getenv(addressEnvKey)
	if addr == "" {
		return defaultAddress
	}
	return addr
}

func getPSQLConnStrFromEnv() string {
	psqlConnStr := os.Getenv(psqlConnStrEnvKey)
	if psqlConnStr == "" {
		return defaultPSQLConnStr
	}
	return psqlConnStr
}
