package main

import (
	"context"
	"grpc_kitchen_microservices/services/common/genproto/orders"
	"log"
	"net/http"
	"time"
	"text/template"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// just creating simple rest apis

type httpServer struct{
	addr string 
}

func NewHttpServer(addr string) *httpServer{
	return &httpServer{addr: addr}
}

func NewGRPCClient(addr string) *grpc.ClientConn{
	conn, err:= grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalf("did not connect: %v", err)
	}
	return conn 
}

func (s *httpServer) Run() error{
	router:= http.NewServeMux()

	conn := NewGRPCClient(":9000")
	defer conn.Close()
	
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c:= orders.NewOrderServiceClient(conn)

		// context with timeout
		ctx, cancel:= context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
	
		_, err:= c.CreateOrder(ctx, &orders.CreateOrderRequest{    // some dummy payload
			CustomerID: 24,
			ProductID: 3123,
			Quantity: 2,
		})
		if err!=nil{
			log.Fatalf("client error: %v", err)
		}

		o, err := c.GetOrders(ctx, &orders.GetOrdersRequest{
			CustomerID: 24,
		})
		if err!=nil{
			log.Fatalf("client error: %v", err)
		}


		t := template.Must(template.New("orders").Parse(ordersTemplate))

		if err := t.Execute(w, o.GetOrders()); err != nil {
			log.Fatalf("template error: %v", err)
		}

	})

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)

}



var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`