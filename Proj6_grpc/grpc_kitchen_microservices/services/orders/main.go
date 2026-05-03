package main

import (
	"log"
)

func main(){
	/*
	The HTTP server runs in a goroutine, and the gRPC server runs on the main thread. 
	Both listen concurrently on ports 8080 and 9000 respectively.
	*/

	httpServer := NewHttpServer(":8080")
	go func() {
		if err := httpServer.Run(); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	grpcServer := NewGRPCServer(":9000")
	if err := grpcServer.Run(); err != nil {
		log.Fatalf("gRPC server error: %v", err)
	}
}