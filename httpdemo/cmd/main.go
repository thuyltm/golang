package main

import (
	"errors"
	"github.com/gorilla/mux"
	"golang-bazel-demo-app/httpdemo/handlers"
	"log"
	"net/http"
	"os"
	"text/template"
)

var templates = template.Must(template.ParseFiles("public/upload.html"))

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/greet", handlers.Greet).Methods("GET")
	router.HandleFunc("/greet-many", handlers.GreetMany).Methods("GET")
	router.HandleFunc("/upload", uploadHandler)
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

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		display(w, "upload", nil)
	case "POST":
		uploadFile(w, r)
	}
}

func display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate
}
