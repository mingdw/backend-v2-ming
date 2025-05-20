package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateComerInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户信息
func NewUpdateComerInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateComerInfoLogic {
	return &UpdateComerInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateComerInfoLogic) UpdateComerInfo(req *types.UpdateComerInfoRequest) (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
