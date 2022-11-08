package srv

import (
	"context"
	"ldm/common/protos/hello"
)

//say hello
func (h *HelloImpl) Hello(ctx context.Context, req *hello.HelloReq, rsp *hello.HelloRsp) error {
	rsp.Msg = "地瓜" + req.Name
	return nil
}
