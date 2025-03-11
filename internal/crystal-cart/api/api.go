package api

import (
	desc "github.com/Dmitriy-M1319/crystal-cart/pkg/crystal-cart/v1"
)

type CartApiImplementation struct {
	desc.UnimplementedCartServiceServer
}

func NewCartApiImplementation() *CartApiImplementation {
	return &CartApiImplementation{}
}
