package main

import (
	"context"
	"fmt"
	"golang-bazel-demo-app/serviceone/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("127.0.0.1:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewServiceOneClient(cc)
	resp, err := client.Hello(context.Background(), &pb.HelloRequest{Name: "Thuy"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("response is: %s\n", resp.Msg)
}
