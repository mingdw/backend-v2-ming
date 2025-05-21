package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateComerLanguagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户语言
func NewUpdateComerLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateComerLanguagesLogic {
	return &UpdateComerLanguagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateComerLanguagesLogic) UpdateComerLanguages() (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
