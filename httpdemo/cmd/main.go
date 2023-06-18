package main

import (
	"errors"
	"github.com/gorilla/mux"
	"golang-bazel-demo-app/httpdemo/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/greet", handlers.Greet).Methods("GET")
	router.HandleFunc("/greet-many", handlers.GreetMany).Methods("GET")
	address := ":5000"
	log.Printf("server started at port %v\n", address)
	err := http.ListenAndServe(address, router)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed \n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
