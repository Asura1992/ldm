package impl

import (
	"go-micro.dev/v4/client"
	"ldm/utils/grpc_err"
)

type ProjectImpl struct {
	client client.Client
}

func NewProjectImpl(cli client.Client)*ProjectImpl{
	return &ProjectImpl{
		client: cli,
	}
}

func CommonGrpcErr(status int,msg string)error{
	return grpc_err.GrpcErr(status,msg)
}
