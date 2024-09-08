package ext

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PlayLineModel = (*customPlayLineModel)(nil)

type (
	// PlayLineModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlayLineModel.
	PlayLineModel interface {
		playLineModel
	}

	customPlayLineModel struct {
		*defaultPlayLineModel
	}
)

// NewPlayLineModel returns a model for the database table.
func NewPlayLineModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PlayLineModel {
	return &customPlayLineModel{
		defaultPlayLineModel: newPlayLineModel(conn, c, opts...),
	}
}
