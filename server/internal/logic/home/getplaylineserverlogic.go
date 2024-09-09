package home

import (
	"context"
	"movies_server/common/client/moviesserver"

	"movies_server/server/internal/svc"
	"movies_server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPlayLineServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPlayLineServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPlayLineServerLogic {
	return &GetPlayLineServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPlayLineServerLogic) GetPlayLineServer(req *types.GetPlayLineRequest) (resp *types.GetPlayLineResponse, err error) {
	// todo: add your logic here and delete this line
	line, err := l.svcCtx.MoviesRpc.GetPlayLine(l.ctx, &moviesserver.GetPlayLineRequest{
		VideoId: req.VideoId,
	})

	resp = &types.GetPlayLineResponse{
		Code:    200,
		Message: "成功",
		Data:    nil,
	}
	if err != nil {
		resp.Code = 0
		resp.Message = err.Error()
	} else {
		data := line.GetData()
		newList := make([]*types.PlayLineModel, 0)
		for _, d := range data {
			p := &types.PlayLineModel{
				Id:           d.GetId(),
				VideoLineId:  d.GetVideoLineId(),
				VideoId:      d.GetVideoId(),
				Name:         d.GetName(),
				Sort:         d.GetSort(),
				File:         d.GetFile(),
				ChargingMode: d.GetChargingMode(),
				Currency:     d.GetCurrency(),
				SubTitle:     d.GetSubTitle(),
				Status:       d.GetStatus(),
				CreateAt:     d.GetCreateAt(),
				UpdateAt:     d.GetUpdateAt(),
				SiteId:       d.GetSiteId(),
				Tag:          d.GetTag(),
				LiveSource:   d.GetLiveSource(),
			}
			newList = append(newList, p)
		}
		resp.Data = newList
	}

	return
}
