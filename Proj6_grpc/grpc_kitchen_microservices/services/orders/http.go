package main

import (
	handler "grpc_kitchen_microservices/services/orders/handler/orders"
	"grpc_kitchen_microservices/services/orders/service"
	"log"
	"net/http"
)

// just creating simple rest apis

type httpServer struct{
	addr string 
}

func NewHttpServer(addr string) *httpServer{
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error{
	router:= http.NewServeMux()

	orderService:= service.NewOrderService()
	orderHandler:= handler.NewHttpOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)

}