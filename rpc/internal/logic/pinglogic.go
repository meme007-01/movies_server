package logic

import (
	"context"

	"movies/rpc/internal/svc"
	"movies/rpc/movies"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *movies.Request) (*movies.Response, error) {
	// todo: add your logic here and delete this line

	return &movies.Response{}, nil
}
