package home

import (
	"context"
	"movies_server/server/internal/client/moviesserver"
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
	// todo: add your logic here and delete this line
	navigation, err := l.svcCtx.MoviesRpc.GetNavigation(l.ctx, &moviesserver.GetNavigationRequest{})
	logx.Infov(navigation)
	if navigation != nil {
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
		resp = &types.GetNavigationResponse{
			Code:    navigation.GetCode(),
			Message: navigation.GetMessage(),
			Data:    list,
		}
	} else {
		resp = &types.GetNavigationResponse{
			Code:    1,
			Message: "失败",
			Data:    nil,
		}
	}
	return resp, err
}
