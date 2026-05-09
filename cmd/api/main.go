package main

import (
	"api-golang/config"
	"api-golang/internal/db"
	"api-golang/internal/handlers"
	"api-golang/internal/handlers/user"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	env := config.LoadEnvs(".env")
	fmt.Printf("DB: %s:%s/%s\n", env.DBHost, env.DBPort, env.DBName)

	pool, err := db.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	r.Get("/healthz", handlers.HealthzHandler)

	r.Route("/user", func(r chi.Router) {
		r.Get("/all", user.GetAllHandler(pool))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
