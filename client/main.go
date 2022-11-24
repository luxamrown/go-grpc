package main

import (
	"context"
	"fmt"
	"log"
	"mohamadelabror/gogrpc/model/user"
	"time"

	"google.golang.org/grpc"
)

const userServiceAddress = "localhost:5001"

func main() {
	// Create connection to gRPC server
	conn, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connec to service: %v", err)
		return
	}

	defer conn.Close()

	userServiceClient := user.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	response, err := userServiceClient.GreetUser(ctx, &user.GreetingRequest{
		Name:       "Marty Mc Fly",
		Salutation: "Mr. ",
	})

	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	// show response
	fmt.Println(response.GreetingMessage)
}
