package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerInvitationCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户邀请人数
func NewGetComerInvitationCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerInvitationCountLogic {
	return &GetComerInvitationCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerInvitationCountLogic) GetComerInvitationCount() (resp *types.ComerInvitationCountResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
