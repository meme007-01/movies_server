package home

import (
	"context"
	"movies_server/common/movies"

	"movies_server/server/internal/svc"
	"movies_server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoServerLogic {
	return &GetVideoServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoServerLogic) GetVideoServer(req *types.GetVideoRequest) (resp *types.GetVideoResponse, err error) {
	// todo: add your logic here and delete this line
	videoData, err := l.svcCtx.MoviesRpc.GetVideoList(l.ctx, &movies.GetVideoRequest{
		PageIndex:       req.PageIndex,
		PageSize:        req.PageSize,
		CategoryPid:     req.CategoryPid,
		CategoryChildId: req.CategoryChildId,
	})

	resp = &types.GetVideoResponse{
		Code:    200,
		Message: "成功",
		Data:    nil,
	}
	if err != nil {
		resp.Code = 0
		resp.Message = err.Error()
	} else {
		data := videoData.GetData()
		list := videoData.GetBannerList()
		videos := make([]*types.MovieModel, 0)
		bannerList := make([]*types.MovieModel, 0)
		for _, m := range data {
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
			bannerList = append(bannerList, movieModel)
		}
		resp.Data = videos
		resp.BannerList = bannerList
		resp.Total = videoData.GetTotal()
	}
	return resp, nil
}
