package main

import (
	"context"
	"log"

	pb "github.com/SebastianFM1/Go-projects/OMSV2/commons/api"
	"google.golang.org/grpc"
)
type grpcHandler struct {
	pb.UnimplementedOrderServiceServer

    service OrdersService
}

func NewGRPCHandler(grpcServer *grpc.Server, service OrdersService) {
	handler := &grpcHandler{
        service: service,
    }
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
    log.Printf("New order received! Order: %v", p)

    o := &pb.Order{
        ID: "42",
        CustomerID: p.CustomerID,
        Status: "CREATED",
        Items: []*pb.Item{}, // De momento vacío, o puedes mapear los ItemsWithQuantity si quieres
    }
    return o, nil
}