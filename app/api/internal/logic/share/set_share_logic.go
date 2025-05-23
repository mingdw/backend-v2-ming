package share

import (
	"context"
	"errors"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comer"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 设置分享
func NewSetShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetShareLogic {
	return &SetShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetShareLogic) SetShare() (resp *types.ShareSetResponse, err error) {
	comerInfo, ok := l.ctx.Value("comerInfo").(*comer.Comer)
	if !ok {
		return nil, errors.New("user not found")
	}
	// 根据用户的address计算分享的cod
	return &types.ShareSetResponse{
		ShareCode: getShareCode(comerInfo.Address),
	}, nil
}

func getShareCode(address string) string {
	if address != "" {
		//获取address的后六位
		return address[len(address)-6:]
	}
	return ""
}
