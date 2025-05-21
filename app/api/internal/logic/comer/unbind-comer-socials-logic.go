package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnbindComerSocialsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 解绑用户社交
func NewUnbindComerSocialsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnbindComerSocialsLogic {
	return &UnbindComerSocialsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnbindComerSocialsLogic) UnbindComerSocials() (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
