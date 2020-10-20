package services

import (
	"context"
	"fmt"
)

type OrderService struct {
}

func (this *OrderService) NewOrder(ctx context.Context, in *OrderRequest) (*OrderResponse, error) {
	err := in.OrderMain.Validate()
	if err != nil {
		return &OrderResponse{Status: "error", Message: err.Error()}, nil
	}
	fmt.Println(">>>", in.OrderMain)
	fmt.Println(">>>", in.OrderMain.OrderDetails)
	return &OrderResponse{Status: "ok", Message: "success"}, nil
}
