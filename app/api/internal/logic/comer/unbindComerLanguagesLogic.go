package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnbindComerLanguagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 解绑用户语言
func NewUnbindComerLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnbindComerLanguagesLogic {
	return &UnbindComerLanguagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnbindComerLanguagesLogic) UnbindComerLanguages() (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
