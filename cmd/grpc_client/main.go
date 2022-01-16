package main

import (
	"context"
	"flag"
	fibonacciService "github.com/Eretic431/fibonacci/internal/fibonacci/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	x := flag.Int("x", 1, "from")
	y := flag.Int("y", 1, "to")
	flag.Parse()

	var conn *grpc.ClientConn
	var opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(":8080", opts...)
	if err != nil {
		log.Fatalf("Could not connect with error: %s", err)
	}
	defer conn.Close()
	c := fibonacciService.NewFibonacciServiceClient(conn)
	response, err := c.Get(context.Background(), &fibonacciService.GetRequest{X: int32(*x), Y: int32(*y)})
	if err != nil {
		log.Fatalf("Error calling get: %s", err)
	}
	log.Println("Response from server: ", response.Numbers)
}
