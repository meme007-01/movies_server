package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"movies_server/rpc/internal/config"
	"movies_server/rpc/models"
)

type ServiceContext struct {
	Config        config.Config
	CategoryModel models.CategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		CategoryModel: models.NewCategoryModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
