package comer

import (
	"context"
	"errors"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"metaLand/data/model/comeraccount"

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
	comerAccountInfo, err := comeraccount.FindComerAccount(l.svcCtx.DB, uint64(req.ComerAccountId))
	if err != nil {
		return nil, err
	}
	if comerAccountInfo == nil {
		return nil, errors.New("comer account not found")
	}
	comerAccountInfo.IsLinked = false
	err = comeraccount.UpdateComerAccount(l.svcCtx.DB, comerAccountInfo)
	if err != nil {
		return nil, err
	}
	return &types.MessageResponse{
		Message: "success",
	}, nil
}
