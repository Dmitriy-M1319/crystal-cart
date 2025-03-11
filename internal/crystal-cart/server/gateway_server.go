package server

import (
	"context"
	"net/http"

	pb "github.com/Dmitriy-M1319/crystal-cart/pkg/crystal-cart/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createGatewayServer(grpcAddr, gatewayAddr string) *http.Server {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterCartServiceHandlerFromEndpoint(ctx, mux, gatewayAddr, opts)
	if err != nil {
		return nil
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Разрешить все источники
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	gatewayServer := &http.Server{
		Addr:    gatewayAddr,
		Handler: handler,
	}

	return gatewayServer
}
