package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"movies_server/rpc/internal/config"
	ext2 "movies_server/rpc/internal/models/ext"
)

type ServiceContext struct {
	Config          config.Config
	ExCategoryModel ext2.ExCategoryModel
	ExVideosModel   ext2.ExVideosModel
	ExPlayLineModel ext2.PlayLineModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ExCategoryModel: ext2.NewExCategoryModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ExVideosModel:   ext2.NewExVideosModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ExPlayLineModel: ext2.NewPlayLineModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
