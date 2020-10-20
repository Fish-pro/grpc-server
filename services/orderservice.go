package services

import (
	"context"
	"fmt"
)

type OrderService struct {
}

func (this *OrderService) NewOrder(ctx context.Context, in *OrderMain) (*OrderResponse, error) {
	fmt.Println(in)
	return &OrderResponse{Status: "ok", Message: "success"}, nil
}
