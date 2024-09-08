package home

import (
	"context"
	"movies_server/common/client/moviesserver"
	"movies_server/server/internal/svc"
	"movies_server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNavigationServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNavigationServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNavigationServerLogic {
	return &GetNavigationServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNavigationServerLogic) GetNavigationServer(req *types.GetNavigationRequest) (resp *types.GetNavigationResponse, err error) {
	navigation, err := l.svcCtx.MoviesRpc.GetNavigation(l.ctx, &moviesserver.GetNavigationRequest{})
	resp = &types.GetNavigationResponse{
		Code:    200,
		Message: "成功",
		Data:    nil,
	}
	if err != nil || navigation == nil {
		resp.Code = 0
		if err != nil {
			resp.Message = err.Error()
		} else {
			resp.Message = "未获取到数据,数据获取为 nil"
		}
	} else {
		data := navigation.GetData()
		list := make([]types.NavigationModel, 0)
		for _, c := range data {
			newList := make([]types.NavigationModel, 0)
			for _, m := range c.SubNavList {
				newList = append(newList, types.NavigationModel{
					Title: m.GetTitle(),
					Id:    m.GetId(),
					Sort:  m.GetSort(),
				})
			}
			list = append(list, types.NavigationModel{
				Title:      c.GetTitle(),
				Id:         c.GetId(),
				Sort:       c.GetSort(),
				SubNavList: newList,
			})
		}
		resp.Data = list
	}
	return resp, err
}
