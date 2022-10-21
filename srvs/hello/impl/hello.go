package impl

import (
	"context"
	"ldm/common/protos/hello"
)
//say hello
func (h HelloImpl) Hello(ctx context.Context, req *hello.HelloReq, rsp *hello.HelloRsp) error {
	rsp.Msg = "hello" + req.Name
	return nil
}
