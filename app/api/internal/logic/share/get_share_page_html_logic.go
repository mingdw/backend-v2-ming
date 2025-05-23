package share

import (
	"context"
	"errors"

	"metaLand/app/api/internal/svc"
	"metaLand/data/model/comer"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSharePageHtmlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取分享
func NewGetSharePageHtmlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSharePageHtmlLogic {
	return &GetSharePageHtmlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSharePageHtmlLogic) GetSharePageHtml() (resp string, err error) {
	// 获取分享码
	comerInfo, ok := l.ctx.Value("comerInfo").(*comer.Comer)
	if !ok {
		return "", errors.New("user not found")
	}

	return comerInfo.Address, nil
}
