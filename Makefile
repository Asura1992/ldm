proto:
	mkdir -p common/swagger && \
	cd common/proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../ --grpc-gateway_out=../ --openapiv2_out=../swagger   hello.proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../ --grpc-gateway_out=../ --openapiv2_out=../swagger   project.proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../ --grpc-gateway_out=../ --openapiv2_out=../swagger   liveroom.proto

h:
	go run ./srvs/hello

p:
	go run ./srvs/project

g:
	go run ./srvs/gateway

l:
	go run ./srvs/liveroom
