package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerInvitationRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户邀请记录
func NewGetComerInvitationRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerInvitationRecordsLogic {
	return &GetComerInvitationRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerInvitationRecordsLogic) GetComerInvitationRecords() (resp *types.PageData, err error) {
	// todo: add your logic here and delete this line

	return
}
