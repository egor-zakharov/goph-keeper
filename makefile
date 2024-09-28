generate_proto:
	protoc \
	--go_out=. --go_opt paths=import \
	--go-grpc_out=. --go-grpc_opt paths=import \
	pkg/proto/gophkeeper/gophkeeper.proto

generate_client_local:
	go build -o client.exe ./cmd/client