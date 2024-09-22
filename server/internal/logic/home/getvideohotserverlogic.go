package home

import (
	"context"
	"movies_server/common/movies"

	"movies_server/server/internal/svc"
	"movies_server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoHotServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoHotServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoHotServerLogic {
	return &GetVideoHotServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoHotServerLogic) GetVideoHotServer(req *types.GetVideoHotRequest) (resp *types.GetVideoHotResponse, err error) {
	// todo: add your logic here and delete this line
	obj, err := l.svcCtx.MoviesRpc.GetVideoHotList(l.ctx, &movies.GetVideoHotRequest{
		CategoryPid: req.CategoryPid,
		TabType:     req.TabType,
	})
	resp = &types.GetVideoHotResponse{
		Code:    200,
		Message: "成功",
		Data:    nil,
	}
	if err != nil {
		resp.Code = 0
		resp.Message = err.Error()
	} else {
		list := obj.GetData()
		videos := make([]*types.MovieModel, 0)
		for _, m := range list {
			movieModel := &types.MovieModel{
				Id:                    m.GetId(),
				Title:                 m.GetTitle(),
				CategoryPid:           m.GetCategoryPid(),
				CategoryChildId:       m.GetCategoryChildId(),
				SurfacePlot:           m.GetSurfacePlot(),
				Recommend:             m.GetRecommend(),
				Cycle:                 m.GetCycle(),
				CycleImg:              m.GetCycleImg(),
				ChargingMode:          m.GetChargingMode(),
				BuyMode:               m.GetBuyMode(),
				Gold:                  m.GetGold(),
				Directors:             m.GetDirectors(),
				Actors:                m.GetActors(),
				ImdbScore:             m.GetImdbScore(),
				ImdbScoreId:           m.GetImdbScoreId(),
				DoubanScore:           m.GetDoubanScore(),
				DoubanScoreId:         m.GetDoubanScoreId(),
				Introduce:             m.GetIntroduce(),
				PopularityDay:         m.GetPopularityDay(),
				PopularityWeek:        m.GetPopularityWeek(),
				PopularityMonth:       m.GetPopularityMonth(),
				PopularitySum:         m.GetPopularitySum(),
				Note:                  m.GetNote(),
				Year:                  m.GetYear(),
				AlbumId:               m.GetAlbumId(),
				Status:                m.GetStatus(),
				CreateAt:              m.GetCreateAt(),
				UpdateAt:              m.GetUpdateAt(),
				Duration:              m.GetDuration(),
				Region:                m.GetRegion(),
				Language:              m.GetLanguage(),
				Label:                 m.GetLabel(),
				Number:                m.GetNumber(),
				Total:                 m.GetTotal(),
				HorizontalPoster:      m.GetHorizontalPoster(),
				VerticalPoster:        m.GetVerticalPoster(),
				Publish:               m.GetPublish(),
				SerialNumber:          m.GetSerialNumber(),
				Screenshot:            m.GetScreenshot(),
				Gif:                   m.GetGif(),
				Alias:                 m.GetAlias(),
				ReleaseAt:             m.GetReleaseAt(),
				ShelfAt:               m.GetShelfAt(),
				End:                   m.GetEnd(),
				Unit:                  m.GetUnit(),
				Watch:                 m.GetWatch(),
				CollectionId:          m.GetCollectionId(),
				UseLocalImage:         m.GetUseLocalImage(),
				TitlesTime:            m.GetTitlesTime(),
				TrailerTime:           m.GetTrailerTime(),
				SiteId:                m.GetSiteId(),
				CategoryPidStatus:     m.GetCategoryPidStatus(),
				CategoryChildIdStatus: m.GetCategoryChildIdStatus(),
				PlayUrl:               m.GetPlayUrl(),
				PlayUrlPutIn:          m.GetPlayUrlPutIn(),
			}
			videos = append(videos, movieModel)
		}
		resp.Data = videos
	}

	return resp, nil
}
