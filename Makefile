proto:
	cd common/proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../   hello.proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../   project.proto

gateway:
	cd common/proto && \
	protoc --proto_path=. --go_out=./gateway/  --go-grpc_out=./gateway/  --grpc-gateway_out=./gateway/   hello.proto && \
    protoc --proto_path=. --go_out=./gateway/  --go-grpc_out=./gateway/  --grpc-gateway_out=./gateway/   project.proto
