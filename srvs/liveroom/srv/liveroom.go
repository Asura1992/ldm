package srv

import (
	"context"
	"ldm/common/protos/liveroom"
)

func (l LiveroomImpl) GetLiveroom(ctx context.Context, req *liveroom.GetLiveroomReq, rsp *liveroom.GetLiveroomRsp) error {
	rsp.Msg = "哈哈哈哈" + req.Name
	return nil
}
