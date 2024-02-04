package main

import (
	"context"
	"fmt"
	pb "github.com/rezaig/grpc-service/story-service/pb/story"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	storyHost = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(storyHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	c := pb.NewStoryServiceClient(conn)

	r, err := c.FindAll(context.Background(), &pb.FindAllRequest{Page: 1, Size: 10})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for _, st := range r.Stories {
		fmt.Println("ID", st.Id)
		fmt.Println("Title", st.Title)
	}
}
