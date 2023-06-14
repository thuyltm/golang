package main

import (
 "fmt"
 "flag"
 "hello_golang/hello"
)

var (
	address         = flag.String("host", "localhost:50051", "host:port of gRPC server")
	skipHealthCheck = flag.Bool("skipHealthCheck", false, "Skip Initial Healthcheck")
)

func main() {
  fmt.Println("hello world")
}
