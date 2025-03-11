package api

import (
	"context"

	pb "github.com/Dmitriy-M1319/crystal-cart/pkg/crystal-cart/v1"
	_ "github.com/Dmitriy-M1319/crystal-services/pkg/crystal-services/products/v1"
)

func (impl *CartApiImplementation) CreateCart(ctx context.Context, email *pb.UserEmail) (*pb.CartIdentifier, error) {
	return nil, nil
}

func (impl *CartApiImplementation) AddProductToCart(ctx context.Context, info *pb.CartProductInfo) (*pb.CartIdentifier, error) {
	return nil, nil
}

func (impl *CartApiImplementation) GetCartProductsByUser(ctx context.Context, email *pb.UserEmail) (*pb.Products, error) {
	return nil, nil
}

func (impl *CartApiImplementation) DeleteProductFromCart(ctx context.Context, info *pb.CartProductInfo) (*pb.CartIdentifier, error) {
	return nil, nil
}
