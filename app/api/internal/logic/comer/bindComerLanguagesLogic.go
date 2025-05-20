package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindComerLanguagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 绑定用户语言
func NewBindComerLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindComerLanguagesLogic {
	return &BindComerLanguagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindComerLanguagesLogic) BindComerLanguages() (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
