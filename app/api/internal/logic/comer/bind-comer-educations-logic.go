package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindComerEducationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 绑定用户教育经历
func NewBindComerEducationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindComerEducationsLogic {
	return &BindComerEducationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindComerEducationsLogic) BindComerEducations() (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
