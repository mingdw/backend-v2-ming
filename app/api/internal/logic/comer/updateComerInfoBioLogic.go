package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateComerInfoBioLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户简介
func NewUpdateComerInfoBioLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateComerInfoBioLogic {
	return &UpdateComerInfoBioLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateComerInfoBioLogic) UpdateComerInfoBio() (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
