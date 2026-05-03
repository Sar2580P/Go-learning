package main

import (
	"log"
)




func main(){
	//  Ports below 1024 often need sudo, but it's better to just use a higher port number
	httpServer := NewHttpServer(":1000")  
	if err := httpServer.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}