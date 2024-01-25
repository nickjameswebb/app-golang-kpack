package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

const defaultPort = "3000"

func name() string {
	n := os.Getenv("CARTO_RUN_WORKLOAD_NAME")
	if n == "" {
		return "golang-web-app"
	}
	return n
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

  fmt.Println("0")

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(fmt.Sprintf("%s is alive", name())))
		if err != nil {
			logrus.Fatalf("unable to write to buffer %v", err)
		}
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK"))
		if err != nil {
			logrus.Fatalf("unable to write to buffer %v", err)
		}
	})

	r.Get("/ready", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(fmt.Sprintf("%s is alive", name())))
		if err != nil {
			logrus.Fatalf("unable to write to buffer %v", err)
		}
	})

	srv := &http.Server{
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           r,
		Addr:              fmt.Sprintf(":%s", port()),
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func port() string {
	s := os.Getenv("PORT")
	if s == "" {
		return defaultPort
	}
	return s
}

func init() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	// LOG_LEVEL not set, let's default to debug
	if !ok {
		lvl = "debug"
	}
	// parse string, this is built-in feature of logrus
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	// set global log level
	logrus.SetLevel(ll)
}
