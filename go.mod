module github.com/Dmitriy-M1319/crystal-cart

go 1.23.0

require (
	github.com/redis/go-redis/v9 v9.7.1
	github.com/rs/cors v1.11.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
)

require (
	github.com/Dmitriy-M1319/crystal-services v0.0.0-20250311055159-8fef2d4668cb
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.25.1
	github.com/rs/zerolog v1.33.0
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto v0.0.0-20250106144421-5f5ef82da422 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250106144421-5f5ef82da422
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250106144421-5f5ef82da422 // indirect
	google.golang.org/grpc v1.69.2
	google.golang.org/protobuf v1.36.2
)

replace github.com/Dmitriy-M1319/crystal-cart/pkg/crystal-cart/v1/ => ./pkg/crystal-cart/v1
