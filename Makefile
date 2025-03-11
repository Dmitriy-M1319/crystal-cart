.PHONY: cart_proto proto_clean cart_run

cart_proto:
	protoc --experimental_allow_proto3_optional \
		-I protos \
		--go_out=pkg --go_opt=paths=source_relative \
		--go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pkg --grpc-gateway_opt=generate_unbound_methods=true --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=swagger \
		protos/crystal-cart/v1/cart.proto
proto_clean:
	rm -rf pkg/crystal-cart
cart_run:
	go run cmd/crystal-cart/main.go
