package main

import (
	"context"
	"log"

	common "github.com/SebastianFM1/Go-projects/OMSV2/commons"
	pb "github.com/SebastianFM1/Go-projects/OMSV2/commons/api"
)

type service struct{
	store OrdersStore
}

func NewService(store OrdersStore) *service{
	return &service{store}
}

func (s *service) CreateOrder(context.Context) error{
	return nil
}

func (s *service) ValidateOrder(ctx context.Context,p *pb.CreateOrderRequest) error{
	if len(p.Items) == 0 {
		return common.ErrNoItems
	}

	mergedItems := mergeItemsQuantity(p.Items)
	log.Print(mergedItems)
	//validate with the stock service

	return nil
}

func mergeItemsQuantity(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity{
	merged := make([]*pb.ItemsWithQuantity,0)

	for _, item := range items {
		found := false
		for _, finalItem := range merged {
			if finalItem.ID == item.ID {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}
		}

		if !found {
			merged= append(merged, item)
		}
	} 

	return merged
}
