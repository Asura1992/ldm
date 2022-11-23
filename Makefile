proto:
	mkdir -p common/swagger && \
	cd common/proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../ --grpc-gateway_out=../ --openapiv2_out=../swagger   hello.proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../ --grpc-gateway_out=../ --openapiv2_out=../swagger   project.proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../ --grpc-gateway_out=../ --openapiv2_out=../swagger   liveroom.proto

#启动网关服务
g:
	go run ./srvs/gateway
#启动hello服务
h:
	go run ./srvs/hello
#启动项目服务
p:
	go run ./srvs/project
#启动房间服务
l:
	go run ./srvs/liveroom
