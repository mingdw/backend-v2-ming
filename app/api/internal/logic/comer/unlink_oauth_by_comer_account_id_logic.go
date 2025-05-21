package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnlinkOauthByComerAccountIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 解绑用户账号
func NewUnlinkOauthByComerAccountIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnlinkOauthByComerAccountIdLogic {
	return &UnlinkOauthByComerAccountIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnlinkOauthByComerAccountIdLogic) UnlinkOauthByComerAccountId(req *types.UnlinkOauthByComerAccountIdRequest) (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
