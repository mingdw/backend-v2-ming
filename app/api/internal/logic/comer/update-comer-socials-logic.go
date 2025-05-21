package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateComerSocialsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户社交
func NewUpdateComerSocialsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateComerSocialsLogic {
	return &UpdateComerSocialsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateComerSocialsLogic) UpdateComerSocials() (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
