package main

import (
	"fmt"
	"log"
	"net/http"

	"api-golang/config"

	"github.com/go-chi/chi/v5"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	r := chi.NewRouter()
	env := config.LoadEnvs(".env")
	fmt.Printf("DB: %s:%s/%s\n", env.DBHost, env.DBPort, env.DBName)

	r.Get("/healthz", handler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
