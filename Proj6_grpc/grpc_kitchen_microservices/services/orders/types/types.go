package types

import (
	"context"
	"grpc_kitchen_microservices/services/common/genproto/orders"
)


type OrderService interface{
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}