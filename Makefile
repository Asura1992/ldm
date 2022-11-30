proto:
	mkdir -p common/swagger && \
	cd common/proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../ --grpc-gateway_out=../ --openapiv2_out=../swagger   hello.proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../ --grpc-gateway_out=../ --openapiv2_out=../swagger   project.proto && \
    protoc --proto_path=. --go_out=../  --go-grpc_out=../  --micro_out=../ --grpc-gateway_out=../ --openapiv2_out=../swagger   liveroom.proto

#网关服务
g:
	docker-compose build --force-rm api-gateway-srv
#hello服务
h:
	docker-compose build --force-rm api-hello-srv
#项目服务
p:
	docker-compose build --force-rm api-project-srv
#房间服务
l:
	docker-compose build --force-rm api-liveroom-srv
#删除无用none镜像 docker rmi $(docker images | grep "none" | awk '{print $3}')
