package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/Dmitriy-M1319/crystal-cart/internal/crystal-cart/api"
	"github.com/Dmitriy-M1319/crystal-services/pkg/crystal-services/products/v1"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	//"github.com/Dmitriy-M1319/crystal-cart/internal/auth/service"
	//"github.com/Dmitriy-M1319/crystal-cart/internal/auth/service/repository"
	"github.com/Dmitriy-M1319/crystal-cart/internal/config"
	pb "github.com/Dmitriy-M1319/crystal-cart/pkg/crystal-cart/v1"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type GrpcServer struct {
	redisConnection *redis.Client
}

func NewGrpcServer(red *redis.Client) *GrpcServer {
	return &GrpcServer{redisConnection: red}
}
func (srv *GrpcServer) Start(conf *config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gatewayAddr := fmt.Sprintf("%s:%v", conf.Grpc.GatewayHost, conf.Grpc.GatewayPort)
	grpcAddr := fmt.Sprintf("%s:%v", conf.Grpc.Host, conf.Grpc.Port)

	gatewayServer := createGatewayServer(grpcAddr, gatewayAddr)

	go func() {
		log.Info().Msgf("Gateway server is running on %s", gatewayAddr)
		if err := gatewayServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Failed running gateway server")
			cancel()
		}
	}()
	isReady := &atomic.Value{}
	isReady.Store(false)

	statusServer := createStatusServer(conf, isReady)

	go func() {
		statusAddr := fmt.Sprintf("%s:%v", conf.Status.Host, conf.Status.Port)
		log.Info().Msgf("Status server is running on %s", statusAddr)
		if err := statusServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Failed running status server")
		}
	}()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer l.Close()

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(conf.Grpc.MaxConnectionIdle) * time.Minute,
			Timeout:           time.Duration(conf.Grpc.Timeout) * time.Second,
			MaxConnectionAge:  time.Duration(conf.Grpc.MaxConnectionAge) * time.Minute,
			Time:              time.Duration(conf.Grpc.Timeout) * time.Minute,
		}),
	)

	// TODO: не забыть про сервисы
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", conf.Product.Host, conf.Product.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to product server: %w", err)
	}
	pb.RegisterCartServiceServer(grpcServer, api.NewCartApiImplementation(products.NewProductsServiceClient(conn)))

	go func() {
		log.Info().Msgf("GRPC Server is listening on: %s", grpcAddr)
		if err := grpcServer.Serve(l); err != nil {
			log.Fatal().Err(err).Msg("Failed running gRPC server")
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		isReady.Store(true)
		log.Info().Msg("The service is ready to accept requests")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Info().Msgf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Info().Msgf("ctx.Done: %v", done)
	}

	isReady.Store(false)

	if err := gatewayServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("gatewayServer.Shutdown")
	} else {
		log.Info().Msg("gatewayServer shut down correctly")
	}

	if err := statusServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("statusServer.Shutdown")
	} else {
		log.Info().Msg("statusServer shut down correctly")
	}

	grpcServer.GracefulStop()
	log.Info().Msgf("grpcServer shut down correctly")

	return nil
}
