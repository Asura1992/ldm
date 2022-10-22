package impl

import (
	"go-micro.dev/v4/client"
	"ldm/utils/grpc_err"
)

type HelloImpl struct {
	client client.Client
}

func NewHelloImplImpl(cli client.Client)*HelloImpl{
	return &HelloImpl{
		client: cli,
	}
}
func CommonGrpcErr(status int,msg string)error{
	return grpc_err.GrpcErr(status,msg)
}
