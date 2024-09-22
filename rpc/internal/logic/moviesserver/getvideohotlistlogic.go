package moviesserverlogic

import (
	"context"
	"movies_server/common/movies"

	"movies_server/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoHotListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoHotListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoHotListLogic {
	return &GetVideoHotListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoHotListLogic) GetVideoHotList(in *movies.GetVideoHotRequest) (*movies.GetVideoHotResponse, error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.ExVideosModel.FindVideoListByHot(l.ctx, in.GetCategoryPid(), in.GetTabType())
	resp := &movies.GetVideoHotResponse{
		Code:    200,
		Message: "成功",
		Data:    nil,
	}
	if err != nil {
		resp.Code = 0
		resp.Message = err.Error()
	} else {
		r := make([]*movies.MovieModel, 0)
		for _, m := range data {
			movie := &movies.MovieModel{
				Id:                    m.Id,
				Title:                 m.Title,
				CategoryPid:           m.CategoryPid,
				CategoryChildId:       m.CategoryChildId,
				SurfacePlot:           m.SurfacePlot.String,
				Recommend:             m.Recommend,
				Cycle:                 m.Cycle,
				CycleImg:              m.CycleImg.String,
				ChargingMode:          m.ChargingMode,
				BuyMode:               m.BuyMode,
				Gold:                  m.Gold,
				Directors:             m.Directors.String,
				Actors:                m.Actors.String,
				ImdbScore:             m.ImdbScore,
				ImdbScoreId:           m.ImdbScoreId,
				DoubanScore:           m.DoubanScore,
				DoubanScoreId:         m.DoubanScoreId,
				Introduce:             m.Introduce.String,
				PopularityDay:         m.PopularityDay,
				PopularityWeek:        m.PopularityWeek,
				PopularityMonth:       m.PopularityMonth,
				PopularitySum:         m.PopularitySum,
				Note:                  m.Note,
				Year:                  m.Year,
				AlbumId:               m.AlbumId,
				Status:                m.Status,
				CreateAt:              m.CreateAt,
				UpdateAt:              m.UpdateAt,
				Duration:              m.Duration,
				Region:                m.Region,
				Language:              m.Language,
				Label:                 m.Label,
				Number:                m.Number.Int64,
				Total:                 m.Total.Int64,
				HorizontalPoster:      m.HorizontalPoster.String,
				VerticalPoster:        m.VerticalPoster.String,
				Publish:               m.Publish.String,
				SerialNumber:          m.SerialNumber.String,
				Screenshot:            m.Screenshot.String,
				Gif:                   m.Gif.String,
				Alias:                 m.Alias.String,
				ReleaseAt:             m.ReleaseAt.Int64,
				ShelfAt:               m.ShelfAt,
				End:                   m.End.Int64,
				Unit:                  m.Unit.String,
				Watch:                 m.Watch.Int64,
				CollectionId:          m.CollectionId.Int64,
				UseLocalImage:         m.UseLocalImage.Int64,
				TitlesTime:            m.TitlesTime,
				TrailerTime:           m.TrailerTime,
				SiteId:                m.SiteId,
				CategoryPidStatus:     m.CategoryPidStatus,
				CategoryChildIdStatus: m.CategoryChildIdStatus,
				PlayUrl:               m.PlayUrl.String,
				PlayUrlPutIn:          m.PlayUrlPutIn,
				TypeSort:              m.TypeSort,
				TypeName:              m.TypeName.String,
			}
			r = append(r, movie)
		}
		resp.Data = r
	}
	return resp, nil
}
