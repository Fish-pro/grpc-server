package services

import (
	"context"
)

type ProdService struct {
}

func (this *ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 20}, nil
}

func (this *ProdService) GetProdStocks(ctx context.Context, in *QuerySize) (*ProdResponseList, error) {
	prodres := []*ProdResponse{
		&ProdResponse{ProdStock: 20},
		&ProdResponse{ProdStock: 21},
		&ProdResponse{ProdStock: 22},
		&ProdResponse{ProdStock: 23},
		&ProdResponse{ProdStock: 24},
		&ProdResponse{ProdStock: 25},
	}
	return &ProdResponseList{Prodres: prodres}, nil
}
