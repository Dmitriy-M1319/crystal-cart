package api

import (
	"github.com/Dmitriy-M1319/crystal-cart/internal/crystal-cart/service"
	desc "github.com/Dmitriy-M1319/crystal-cart/pkg/crystal-cart/v1"
	"github.com/Dmitriy-M1319/crystal-services/pkg/crystal-services/products/v1"
)

type CartApiImplementation struct {
	desc.UnimplementedCartServiceServer
	productClient products.ProductsServiceClient
	service       *service.CartService
}

func NewCartApiImplementation(pClient products.ProductsServiceClient, srv *service.CartService) *CartApiImplementation {
	return &CartApiImplementation{productClient: pClient, service: srv}
}
