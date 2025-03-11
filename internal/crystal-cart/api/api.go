package api

import (
	desc "github.com/Dmitriy-M1319/crystal-cart/pkg/crystal-cart/v1"
	"github.com/Dmitriy-M1319/crystal-services/pkg/crystal-services/products/v1"
)

type CartApiImplementation struct {
	desc.UnimplementedCartServiceServer
	productClient products.ProductsServiceClient
}

func NewCartApiImplementation(pClient products.ProductsServiceClient) *CartApiImplementation {
	return &CartApiImplementation{productClient: pClient}
}
