package comer

import (
	"context"
	"errors"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comer"

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
	comerInfo, ok := l.ctx.Value("comerInfo").(*comer.Comer)
	if !ok {
		return nil, errors.New("user not found")
	}
	l.Logger.Infof("comerInfo: %+v", comerInfo)
	return
}
