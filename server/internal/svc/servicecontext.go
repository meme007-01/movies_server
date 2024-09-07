package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"movies_server/server/internal/client/moviesserver"
	"movies_server/server/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	MoviesRpc moviesserver.MoviesServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		MoviesRpc: moviesserver.NewMoviesServer(zrpc.MustNewClient(c.MoviesRpc)),
	}
}
