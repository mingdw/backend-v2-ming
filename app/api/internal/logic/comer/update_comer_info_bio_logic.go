package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comerprofile"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateComerInfoBioLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户简介
func NewUpdateComerInfoBioLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateComerInfoBioLogic {
	return &UpdateComerInfoBioLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateComerInfoBioLogic) UpdateComerInfoBio(req *types.UpdateComerInfoRequest) (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line
	comerInfo, err := comerprofile.FindComerProfile(l.svcCtx.DB, uint64(req.ComerId))
	if err != nil {
		return nil, err
	}
	comerInfo.Bio = req.Bio
	err = comerprofile.UpdateComerProfile(l.svcCtx.DB, uint64(req.ComerId), map[string]interface{}{
		"bio": req.Bio,
	})
	if err != nil {
		return nil, err
	}
	return &types.MessageResponse{
		Message: "success",
	}, nil
}
