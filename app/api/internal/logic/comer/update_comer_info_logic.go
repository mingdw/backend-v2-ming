package comer

import (
	"context"
	"errors"
	"time"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comer"
	"metaLand/data/model/comerprofile"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateComerInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户信息
func NewUpdateComerInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateComerInfoLogic {
	return &UpdateComerInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateComerInfoLogic) UpdateComerInfo(req *types.UpdateComerInfoRequest) (resp *types.MessageResponse, err error) {
	if req.ComerId == 0 {
		return nil, errors.New("comer_id is required")
	}

	comer, err := comer.FindComer(l.svcCtx.DB, uint64(req.ComerId))
	if err != nil {
		return nil, err
	}
	if comer == nil {
		return nil, errors.New("comer not found")
	}

	comerProfileInfo, err := comerprofile.FindComerProfile(l.svcCtx.DB, uint64(req.ComerId))
	if err != nil {
		return nil, err
	}

	// 如果不存在则创建新的 profile
	if comerProfileInfo == nil {
		comerProfileInfo = &comerprofile.ComerProfile{
			ComerId:    int64(req.ComerId),
			Name:       req.Name,
			Avatar:     req.Avatar,
			Cover:      req.Cover,
			Location:   req.Location,
			TimeZone:   req.TimeZone,
			Website:    req.Website,
			Email:      req.Email,
			Twitter:    req.Twitter,
			Discord:    req.Discord,
			Telegram:   req.Telegram,
			Medium:     req.Medium,
			Facebook:   req.Facebook,
			Linktree:   req.Linktree,
			Bio:        req.Bio,
			Languages:  req.Languages,
			Educations: req.Educations,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			IsDeleted:  0,
		}
		err = comerprofile.InsertComerProfile(l.svcCtx.DB, comerProfileInfo)
	} else {
		// 更新现有 profile
		updates := map[string]interface{}{
			"name":       req.Name,
			"avatar":     req.Avatar,
			"cover":      req.Cover,
			"location":   req.Location,
			"time_zone":  req.TimeZone,
			"website":    req.Website,
			"email":      req.Email,
			"twitter":    req.Twitter,
			"discord":    req.Discord,
			"telegram":   req.Telegram,
			"medium":     req.Medium,
			"facebook":   req.Facebook,
			"linktree":   req.Linktree,
			"bio":        req.Bio,
			"languages":  req.Languages,
			"educations": req.Educations,
			"updated_at": time.Now(),
		}
		err = comerprofile.UpdateComerProfile(l.svcCtx.DB, uint64(req.ComerId), updates)
	}

	if err != nil {
		logx.Errorf("update comer profile failed: %v", err)
		return nil, err
	}

	return &types.MessageResponse{
		Message: "success",
	}, nil
}
