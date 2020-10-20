package services

import (
	"context"
)

type ProdService struct {
}

func (this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0
	switch request.ProdArea {
	case ProdAreas_A:
		stock = 10
	case ProdAreas_B:
		stock = 20
	case ProdAreas_C:
		stock = 30
	default:
		stock = 50
	}
	return &ProdResponse{ProdStock: stock}, nil
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

func (this *ProdService) GetProdInfo(ctx context.Context, in *ProdRequest) (*ProdModel, error) {
	ret := ProdModel{
		ProdId:    101,
		ProdName:  "测试商品",
		ProdPrice: 20.5,
	}
	return &ret, nil
}
