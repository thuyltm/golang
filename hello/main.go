package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang-bazel-demo-app/hello/room"
)

var (
	address         = flag.String("host", "localhost:50051", "host:port of gRPC server")
	skipHealthCheck = flag.Bool("skipHealthCheck", false, "Skip Initial Healthcheck")
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	port := "8080"
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	fmt.Fprintln(w, "Hello, world!")
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hostname: %s\n", host)
	room.PrintDetails(112, 3, 2)
	i := 1
	fmt.Fprintf(w, "initial: %d\n", i)
	room.Zeroval(i)
	fmt.Fprintf(w, "zeroval: %d\n", i)
	room.Zeroptr(&i)
	fmt.Fprintf(w, "zeroptr: %d\n", i)
	fmt.Fprintf(w, "pointer: %p\n", &i)
	cat := room.Cat{Color: "blue", Age: 8, Name: "Milow"}
	cat.Rename("Bob")
	fmt.Fprintln(w, cat.Name)
	cat.RenameV2("Ben")
	fmt.Fprintln(w, cat.Name)
}
