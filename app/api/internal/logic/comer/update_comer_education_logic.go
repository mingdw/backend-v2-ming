package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateComerEducationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户教育经历
func NewUpdateComerEducationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateComerEducationLogic {
	return &UpdateComerEducationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateComerEducationLogic) UpdateComerEducation() (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
