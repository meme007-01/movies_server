package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"movies_server/rpc/internal/config"
	"movies_server/rpc/models/ext"
)

type ServiceContext struct {
	Config          config.Config
	ExCategoryModel ext.ExCategoryModel
	ExVideosModel   ext.ExVideosModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ExCategoryModel: ext.NewExCategoryModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ExVideosModel:   ext.NewExVideosModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
