package handler

import (
	"context"
	"grpc_kitchen_microservices/services/common/genproto/orders"
	"grpc_kitchen_microservices/services/orders/types"

	"google.golang.org/grpc"
)

type OrderGrpcHandler struct{
	// service injection
	orderService types.OrderService

	// 🌟 embedding an unimplemented server stub
	orders.UnimplementedOrderServiceServer
}


func NewGrpcOrdersService(grpc *grpc.Server, orderService types.OrderService){
	gRPCHandler:= &OrderGrpcHandler{
		orderService: orderService,
	}

	// register OrderServiceServer
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrderGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrderResponse, error) {
	o := h.orderService.GetOrders(ctx)
	res := &orders.GetOrderResponse{
		Orders: o,
	}

	return res, nil
}

func (h *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {

	// some mock data
	order:= &orders.Order{
		OrderID: 42,
		CustomerID: 2,
		ProductID: 1,
		Quantity: 10,
	}

	err:= h.orderService.CreateOrder(ctx, order)
	if err!=nil{
		return nil, err
	}

	res:= &orders.CreateOrderResponse{
		Status: "success",
	}
	return res, err
}