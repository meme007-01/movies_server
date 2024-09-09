package moviesserverlogic

import (
	"context"
	"movies_server/common/movies"

	"movies_server/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPlayLineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPlayLineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPlayLineLogic {
	return &GetPlayLineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPlayLineLogic) GetPlayLine(in *movies.GetPlayLineRequest) (*movies.GetPlayLineResponse, error) {
	list, err := l.svcCtx.ExPlayLineModel.FindList(l.ctx, in.GetVideoId())

	resp := &movies.GetPlayLineResponse{
		Code:    200,
		Message: "成功",
		Data:    nil,
	}
	if err != nil {
		resp.Code = 0
		resp.Message = err.Error()
	} else {
		newList := make([]*movies.PlayLineModel, 0)
		for _, line := range list {
			playLine := &movies.PlayLineModel{
				Id:           line.Id,
				VideoLineId:  line.VideoLineId.Int64,
				VideoId:      line.VideoId.Int64,
				Name:         line.Name,
				Sort:         line.Sort,
				File:         line.File.String,
				ChargingMode: line.ChargingMode,
				Currency:     line.Currency,
				SubTitle:     line.SubTitle,
				Status:       line.Status,
				CreateAt:     line.CreateAt,
				UpdateAt:     line.UpdateAt,
				SiteId:       line.SiteId,
				Tag:          line.Tag,
				LiveSource:   line.LiveSource,
			}
			newList = append(newList, playLine)
		}
		resp.Data = newList
	}
	return resp, nil
}
