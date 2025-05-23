package comer

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comereducation"

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

func (l *UpdateComerEducationLogic) UpdateComerEducation(req *types.UpdateComerEducationRequest) (resp *types.MessageResponse, err error) {
	comerEducation := comereducation.ComerEducation{
		Id:          uint64(req.ComerEducationId),
		ComerId:     uint64(req.ComerId),
		School:      req.School,
		Degree:      req.Degree,
		Major:       req.Major,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Description: req.Description,
	}
	err = comereducation.UpdateComerEducation(l.svcCtx.DB, &comerEducation)
	if err != nil {
		return nil, err
	}

	return &types.MessageResponse{
		Message: "success",
	}, nil
}
