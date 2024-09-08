package ext

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"movies_server/rpc/internal/models"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	categoryFieldNames          = builder.RawFieldNames(&Category{})
	categoryRows                = strings.Join(categoryFieldNames, ",")
	categoryRowsExpectAutoSet   = strings.Join(stringx.Remove(categoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	categoryRowsWithPlaceHolder = strings.Join(stringx.Remove(categoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheCmsCategoryIdPrefix = "cache:cms:category:id:"
)

type (
	categoryModel interface {
		Insert(ctx context.Context, data *Category) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Category, error)
		Update(ctx context.Context, data *Category) error
		Delete(ctx context.Context, id int64) error
		FindList(ctx context.Context) ([]Category, error)
	}

	defaultCategoryModel struct {
		sqlc.CachedConn
		table string
	}

	Category struct {
		Id         int64  `db:"id"`
		ParentId   int64  `db:"parent_id"` // 父id
		Type       int64  `db:"type"`      // 类型 1影片 2名人 3文章
		Name       string `db:"name"`      // 分类名称
		Sort       int64  `db:"sort"`      // 排序
		CreateAt   int64  `db:"create_at"`
		UpdateAt   int64  `db:"update_at"`
		IsVertical int64  `db:"is_vertical"` // 是否是竖屏，1-是，0-否
		IsFont     int64  `db:"is_font"`     // 是否是纯文字，1-是，0-否
		SiteId     int64  `db:"site_id"`     // 站点id
		Status     int64  `db:"status"`
	}
)

func newCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultCategoryModel {
	return &defaultCategoryModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`category`",
	}
}

func (m *defaultCategoryModel) Delete(ctx context.Context, id int64) error {
	cmsCategoryIdKey := fmt.Sprintf("%s%v", cacheCmsCategoryIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, cmsCategoryIdKey)
	return err
}

func (m *defaultCategoryModel) FindOne(ctx context.Context, id int64) (*Category, error) {
	cmsCategoryIdKey := fmt.Sprintf("%s%v", cacheCmsCategoryIdPrefix, id)
	var resp Category
	err := m.QueryRowCtx(ctx, &resp, cmsCategoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", categoryRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, models.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCategoryModel) Insert(ctx context.Context, data *Category) (sql.Result, error) {
	cmsCategoryIdKey := fmt.Sprintf("%s%v", cacheCmsCategoryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, categoryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ParentId, data.Type, data.Name, data.Sort, data.IsVertical, data.IsFont, data.SiteId, data.Status)
	}, cmsCategoryIdKey)
	return ret, err
}

func (m *defaultCategoryModel) Update(ctx context.Context, data *Category) error {
	cmsCategoryIdKey := fmt.Sprintf("%s%v", cacheCmsCategoryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, categoryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ParentId, data.Type, data.Name, data.Sort, data.IsVertical, data.IsFont, data.SiteId, data.Status, data.Id)
	}, cmsCategoryIdKey)
	return err
}

func (m *defaultCategoryModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheCmsCategoryIdPrefix, primary)
}

func (m *defaultCategoryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", categoryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCategoryModel) tableName() string {
	return m.table
}

// FindList 查询所有的大分类
func (m *defaultCategoryModel) FindList(ctx context.Context) ([]Category, error) {
	query := fmt.Sprintf("select * from %s where parent_id=0 and status=1 and type=1;", m.table)
	var resp []Category
	allCmsCategoryIdKey := "cache:cms:videos:category"
	err := m.GetCacheCtx(ctx, allCmsCategoryIdKey, &resp)
	if err != nil || resp == nil {
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query)
		if err == nil && resp != nil {
			err1 := m.SetCacheWithExpireCtx(ctx, allCmsCategoryIdKey, resp, time.Hour*12)
			if err1 != nil {
				logx.Error("设置缓存失败了:", err1)
			}
		}
	}
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, models.ErrNotFound
	default:
		return nil, err
	}
}
