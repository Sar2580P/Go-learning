package service

import (
	"context"
	"grpc_kitchen_microservices/services/common/genproto/orders"
	"grpc_kitchen_microservices/services/orders/types"
)
var _ types.OrderService=(*OrderService)(nil)

// holds business logic, separate from transport, db, tcp etc...

var ordersDB =make([]*orders.Order, 0)

type OrderService struct{
	// store(dependency injection)
}

func NewOrderService() *OrderService{
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDB=append(ordersDB, order) 
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return ordersDB
}
