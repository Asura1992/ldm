proto:
	mkdir -p common/swagger && \
	cd common/proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../ --grpc-gateway_out=../ --openapiv2_out=../swagger   hello.proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../ --grpc-gateway_out=../ --openapiv2_out=../swagger   project.proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../ --grpc-gateway_out=../ --openapiv2_out=../swagger   liveroom.proto

#启动网关服务
g:
	docker-compose build --force-rm api-gateway-srv
#启动hello服务
h:
	docker-compose build --force-rm api-hello-srv
#启动项目服务
p:
	docker-compose build --force-rm api-project-srv
#启动房间服务
l:
	docker-compose build --force-rm api-liveroom-srv

up:
	docker-compose up
