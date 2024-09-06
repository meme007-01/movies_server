package moviesserverlogic

import (
	"context"
	"movies_server/rpc/internal/svc"
	"movies_server/rpc/movies"
	"movies_server/rpc/moviesserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNavigationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNavigationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNavigationLogic {
	return &GetNavigationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNavigationLogic) GetNavigation(in *movies.GetNavigationRequest) (*movies.GetNavigationResponse, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.CategoryModel.FindList(l.ctx)
	logx.Infov(list)
	res := &movies.GetNavigationResponse{
		Code:    200,
		Message: "成功",
	}
	if list != nil {
		newList := make([]*moviesserver.NavigationModel, 0)
		for _, category := range list {
			if category.ParentId == 0 && category.Type == 1 {
				subLIst := make([]*moviesserver.NavigationModel, 0)

				for _, c := range list {
					if c.ParentId == category.Id {
						subLIst = append(subLIst, &moviesserver.NavigationModel{
							Title: c.Name,
							Id:    c.Id,
						})
					}
				}
				nav := &moviesserver.NavigationModel{
					Id:         category.Id,
					Title:      category.Name,
					SubNavList: subLIst,
				}
				newList = append(newList, nav)
			}
		}
		res.Data = newList
	} else {
		res.Code = -1
		res.Message = err.Error()
		res.Data = nil
	}
	return res, nil
}
