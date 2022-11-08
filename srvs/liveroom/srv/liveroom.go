package srv

import (
	"context"
	"fmt"
	"ldm/common/protos/liveroom"
)

func (l LiveroomImpl) GetLiveroom(ctx context.Context, req *liveroom.GetLiveroomReq, rsp *liveroom.GetLiveroomRsp) error {
	rsp.Msg = "哈哈哈哈" + req.Name
	fmt.Println("11111111111111")
	return nil
}
