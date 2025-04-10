package api

import (
	"context"

	pb "github.com/Dmitriy-M1319/crystal-cart/pkg/crystal-cart/v1"
	prodPb "github.com/Dmitriy-M1319/crystal-services/pkg/crystal-services/products/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (impl *CartApiImplementation) CreateCart(ctx context.Context, email *pb.UserEmail) (*pb.Empty, error) {
	// старт трассы, сбор метрик(успешных/неудачных запросов)
	err := impl.service.CreateCart(email.Email)
	if err != nil {
		// неудачное окончание трассы и метрики
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.Empty{}, nil
}

func (impl *CartApiImplementation) AddProductToCart(ctx context.Context, info *pb.CartProductInfo) (*pb.Empty, error) {
	// старт трассы, сбор метрик(успешных/неудачных запросов)
	err := impl.service.AddProduct(info.UserEmail, info.ProductId, info.Count)
	if err != nil {
		// неудачное окончание трассы и метрики
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.Empty{}, nil
}

func (impl *CartApiImplementation) GetCartProductsByUser(ctx context.Context, email *pb.UserEmail) (*pb.Products, error) {
	// старт трассы, сбор метрик(успешных/неудачных запросов)
	products, err := impl.service.GetProducts(email.Email)
	if err != nil {
		// неудачное окончание трассы и метрики
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	protoProducts := pb.Products{Products: make([]*pb.ProductInfo, len(products))}
	for _, p := range products {

		externalProduct, err := impl.productClient.GetProductById(ctx, &prodPb.ProductId{Id: uint64(p.ProductID)})
		if err != nil {
			// stat, ok := status.FromError(err)
			// TODO: make correct status returning in product service
			return nil, status.Error(codes.Aborted, err.Error())
		}
		protoProducts.Products = append(protoProducts.Products, &pb.ProductInfo{Product: externalProduct, Count: p.Count})
	}
	return &protoProducts, nil
}

func (impl *CartApiImplementation) DeleteProductFromCart(ctx context.Context, info *pb.CartProductInfo) (*pb.Empty, error) {
	// старт трассы, сбор метрик(успешных/неудачных запросов)
	err := impl.service.RemoveProduct(info.UserEmail, info.ProductId)
	if err != nil {
		// неудачное окончание трассы и метрики
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.Empty{}, nil
}
