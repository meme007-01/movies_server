package home

import (
	"context"
	"movies_server/server/internal/movies"

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
	// todo: add your logic here and delete this line
	recommend, err := l.svcCtx.MoviesRpc.GetRecommend(l.ctx, &movies.GetRecommendRequest{})
	data := recommend.GetData()
	recommendList := make([]*types.RecommendList, 0)
	bannerRecommend := &types.RecommendList{
		BannerList: make([]*types.MovieModel, 0),
		List:       nil,
		Type:       0,
		Name:       "幻灯片",
	}

	dataMap := map[int64]*types.RecommendList{}

	typeMap := map[int64]string{1: "电影", 2: "电视剧", 3: "综艺", 4: "动漫", 39: "爽文短剧", 52: "预告片", 53: "伦理"}
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
			_, strOk := typeMap[typeId]
			if !strOk {
				continue
			}
			_, ok := dataMap[typeId]
			if !ok {
				dataMap[typeId] = &types.RecommendList{
					BannerList: make([]*types.MovieModel, 0),
					List:       make([]*types.MovieModel, 0),
					Type:       typeId,
					Name:       typeMap[typeId],
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
	return &types.GetRecommendResponse{
		Code:    200,
		Message: "成功",
		Data:    recommendList,
	}, nil
}
