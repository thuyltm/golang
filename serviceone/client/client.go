package client

import (
	"context"
	"fmt"
	"golang-bazel-demo-app/serviceone/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("127.0.0.1:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.ServiceOneClient(cc)
	resp, err := client.Hello(context.Background(), &pb.Request{name: ""})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("response is: %d\n", resp.C)
}
