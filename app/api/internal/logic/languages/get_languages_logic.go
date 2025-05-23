package languages

import (
	"context"
	"errors"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comer"
	"metaLand/data/model/comerlanguage"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLanguagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取语言列表
func NewGetLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguagesLogic {
	return &GetLanguagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLanguagesLogic) GetLanguages() (resp *types.LanguageResponse, err error) {
	comerInfo, ok := l.ctx.Value("comerInfo").(*comer.Comer)
	if !ok {
		return nil, errors.New("user not found")
	}
	comerLanguages, err := comerlanguage.ListComerLanguages(l.svcCtx.DB, uint64(comerInfo.Id))
	if err != nil {
		return nil, err
	}
	res := make([]types.LanguageResponse, 0, len(comerLanguages))
	for _, comerLanguage := range comerLanguages {
		res = append(res, types.LanguageResponse{
			Id:       int(comerLanguage.Id),
			ComerId:  int(comerLanguage.ComerId),
			Language: comerLanguage.Language,
			Code:     comerLanguage.Language,
			Level:    comerLanguage.Level,
			IsNative: comerLanguage.IsNative,
		})
	}

	return resp, nil
}
