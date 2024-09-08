package home

import (
	"context"
	"movies_server/common/movies"

	"movies_server/server/internal/svc"
	"movies_server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecommendServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecommendServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecommendServerLogic {
	return &GetRecommendServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecommendServerLogic) GetRecommendServer(req *types.GetRecommendRequest) (resp *types.GetRecommendResponse, err error) {
	r, err1 := l.svcCtx.MoviesRpc.GetRecommend(l.ctx, &movies.GetRecommendRequest{})
	resp = &types.GetRecommendResponse{
		Code:    200,
		Message: "成功",
		Data:    nil,
	}
	if err1 != nil || r == nil {
		resp.Code = 0
		if err1 != nil {
			resp.Message = err1.Error()
		} else {
			resp.Message = "没有获取到数据,数据是 nil"
		}
	} else {
		data := r.GetData()
		recommendList := make([]*types.RecommendList, 0)
		bannerRecommend := &types.RecommendList{
			BannerList: make([]*types.MovieModel, 0),
			List:       nil,
			Type:       0,
			Name:       "幻灯片",
			Sort:       0,
		}
		dataMap := map[int64]*types.RecommendList{}
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
			if movieModel.Cycle == 1 {
				bannerRecommend.BannerList = append(bannerRecommend.BannerList, movieModel)
			} else {
				typeId := m.GetCategoryPid()
				//_, strOk := typeMap[typeId]
				//if !strOk {
				//	continue
				//}
				_, ok := dataMap[typeId]
				if !ok {
					dataMap[typeId] = &types.RecommendList{
						BannerList: make([]*types.MovieModel, 0),
						List:       make([]*types.MovieModel, 0),
						Type:       typeId,
						Name:       m.TypeName,
						Sort:       m.TypeSort,
					}
				}
				o, _ := dataMap[typeId]
				o.List = append(o.List, movieModel)
			}
		}
		recommendList = append(recommendList, bannerRecommend)
		for _, list := range dataMap {
			recommendList = append(recommendList, list)
		}
		resp.Data = recommendList
	}
	return resp, nil
}
