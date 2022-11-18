package srv

import (
	"context"
	"fmt"
	"ldm/common/protos/hello"
	"time"
)

//say hello
func (h *HelloImpl) Hello(ctx context.Context, req *hello.HelloReq, rsp *hello.HelloRsp) error {
	rsp.Msg = "地瓜" + req.Name
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	return nil
}
