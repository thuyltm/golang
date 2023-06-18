package main

import (
	"flag"
	"fmt"

	"golang-bazel-demo-app/hello/room"
)

var (
	address         = flag.String("host", "localhost:50051", "host:port of gRPC server")
	skipHealthCheck = flag.Bool("skipHealthCheck", false, "Skip Initial Healthcheck")
)

func main() {
	fmt.Println("hello world")
	room.PrintDetails(112, 3, 2)
	i := 1
	fmt.Println("initial:", i)
	room.Zeroval(i)
	fmt.Println("zeroval:", i)
	room.Zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
	cat := room.Cat{Color: "blue", Age: 8, Name: "Milow"}
	cat.Rename("Bob")
	fmt.Println(cat.Name)
	cat.RenameV2("Ben")
	fmt.Println(cat.Name)
}
