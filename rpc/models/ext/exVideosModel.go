package ext

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ExVideosModel = (*customVideosModel)(nil)

type (
	// ExVideosModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideosModel.
	ExVideosModel interface {
		videosModel
	}

	customVideosModel struct {
		*defaultVideosModel
	}
)

// NewExVideosModel returns a model for the database table.
func NewExVideosModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ExVideosModel {
	return &customVideosModel{
		defaultVideosModel: newVideosModel(conn, c, opts...),
	}
}
