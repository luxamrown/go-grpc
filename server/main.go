package main

import (
	"context"
	"fmt"
	"log"
	"mohamadelabror/gogrpc/model/user"
	"net"

	"google.golang.org/grpc"
)

//userService is a struct implementing UserServiceServer
type userService struct{}

// GreetUser return greeting message given the name and salutation in gRPC protocol
func (*userService) GreetUser(ctx context.Context, req *user.GreetingRequest) (*user.GreetingResponse, error) {
	// business logic
	salutationMessage := fmt.Sprintf("Howdy, %s %s, nice to see you in the future!", req.Salutation, req.Name)
	return &user.GreetingResponse{GreetingMessage: salutationMessage}, nil
}

func main() {
	// create tcp server on localhost:5001
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}

	// create new gRPC server handler
	server := grpc.NewServer()
	user.RegisterUserServiceServer(server, &userService{})

	// run server

	if err := server.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
