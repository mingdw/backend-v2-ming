package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerInfoDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户详情
func NewGetComerInfoDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerInfoDetailLogic {
	return &GetComerInfoDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerInfoDetailLogic) GetComerInfoDetail() (resp *types.ComerInfoDetailResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
