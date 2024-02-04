package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/rezaig/grpc-service/story-service/internal/delivery/grpcsvc"
	pb "github.com/rezaig/grpc-service/story-service/pb/story"
	"google.golang.org/grpc"
)

const (
	port = 50051
)

func main() {
	sigCh := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	quitCh := make(chan bool, 1)
	signal.Notify(sigCh, os.Interrupt)

	// Graceful shutdown
	go func() {
		for {
			select {
			case <-sigCh:
				log.Println("shutdown due to interrupt signal")
				quitCh <- true
			case e := <-errCh:
				log.Printf("shutdown due to error, error: %v\n", e)
				quitCh <- true
			}
		}
	}()

	// Start gRPC server
	go func() {
		s := grpc.NewServer()
		svc := grpcsvc.NewService()
		pb.RegisterStoryServiceServer(s, svc)
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			errCh <- err
			return
		}
		log.Printf("listening on %d", port)
		errCh <- s.Serve(lis)
	}()

	<-quitCh
	log.Printf("exiting")
}
