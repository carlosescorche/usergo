package main

import (
	"log"
	"net/http"
	"os"

	"github.com/carlosescorche/usergo/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", handlers.HandlerUserAdd).Methods("POST")
	r.HandleFunc("/user/{id}", handlers.HandlerUserUpdate).Methods("PUT")

	err := http.ListenAndServe(os.Getenv("HTTP_PORT"), r)

	if err != nil {
		log.Fatal(err.Error())
	}
}
