package ext

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ExCategoryModel = (*customCategoryModel)(nil)

type (
	// ExCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCategoryModel.
	ExCategoryModel interface {
		categoryModel
	}

	customCategoryModel struct {
		*defaultCategoryModel
	}
)

// NewExCategoryModel returns a model for the database table.
func NewExCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ExCategoryModel {
	return &customCategoryModel{
		defaultCategoryModel: newCategoryModel(conn, c, opts...),
	}
}
