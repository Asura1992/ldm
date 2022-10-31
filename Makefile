proto:
	cd common/proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../   hello.proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../   project.proto

gateway:
	cd common/proto && \
	protoc --proto_path=. --go_out=.  --go-grpc_out=.  --grpc-gateway_out=. --openapiv2_out=../swagger   hello.proto && \
    protoc --proto_path=. --go_out=.  --go-grpc_out=.  --grpc-gateway_out=. --openapiv2_out=../swagger   project.proto
h:
	go run ./srvs/hello

p:
	go run ./srvs/project

g:
	go run ./srvs/gateway
