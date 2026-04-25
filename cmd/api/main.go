package main

import (
	"fmt"
	"log"
	"net/http"

	"api-golang/config"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método nao aceito", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	env := config.LoadEnvs(".env")
	fmt.Printf("DB: %s:%s/%s\n", env.DBHost, env.DBPort, env.DBName)

	http.HandleFunc("/healthz", handler)

	addr := ":" + env.AppPort
	fmt.Printf("Servidor rodando na porta %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("erro ao iniciar o servidor:", err)
	}
}
