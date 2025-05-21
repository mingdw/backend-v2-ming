package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnbindComerEducationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 解绑用户教育经历
func NewUnbindComerEducationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnbindComerEducationsLogic {
	return &UnbindComerEducationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnbindComerEducationsLogic) UnbindComerEducations() (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
