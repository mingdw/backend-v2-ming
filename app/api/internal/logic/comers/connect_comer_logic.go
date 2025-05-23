package comers

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConnectComerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 连接某个comer
func NewConnectComerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConnectComerLogic {
	return &ConnectComerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConnectComerLogic) ConnectComer(req *types.ConnectComerRequest) (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
