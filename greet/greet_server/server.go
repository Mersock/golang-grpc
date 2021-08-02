package main

import (
	"fmt"
	"log"
	"net"

	"golang-grpc/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	fmt.Println("hello")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Fail to listen %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fail to server %v", err)
	}
}
