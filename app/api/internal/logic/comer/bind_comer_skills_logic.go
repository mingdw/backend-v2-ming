package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindComerSkillsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 绑定用户技能
func NewBindComerSkillsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindComerSkillsLogic {
	return &BindComerSkillsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindComerSkillsLogic) BindComerSkills() (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
