package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindComerSocialsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 绑定用户社交
func NewBindComerSocialsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindComerSocialsLogic {
	return &BindComerSocialsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindComerSocialsLogic) BindComerSocials() (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
