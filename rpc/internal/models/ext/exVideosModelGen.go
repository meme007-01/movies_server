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
	videosFieldNames          = builder.RawFieldNames(&Videos{})
	videosRows                = strings.Join(videosFieldNames, ",")
	videosRowsExpectAutoSet   = strings.Join(stringx.Remove(videosFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	videosRowsWithPlaceHolder = strings.Join(stringx.Remove(videosFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheCmsVideosIdPrefix                                    = "cache:cms:videos:id:"
	cacheCmsVideosSiteIdCategoryPidCategoryChildIdTitlePrefix = "cache:cms:videos:siteId:categoryPid:categoryChildId:title:"
)

type (
	videosModel interface {
		Insert(ctx context.Context, data *Videos) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Videos, error)
		FindOneBySiteIdCategoryPidCategoryChildIdTitle(ctx context.Context, siteId int64, categoryPid int64, categoryChildId int64, title string) (*Videos, error)
		Update(ctx context.Context, data *Videos) error
		Delete(ctx context.Context, id int64) error

		FindRecommendList(ctx context.Context) (data []*Videos, err error)
		FindBannerList(ctx context.Context, categoryPid int64) (data []*Videos, err error)
		FindVideoList(ctx context.Context, pageIndex, pageSize, categoryPid, categoryChildId int64) (data []*Videos, err error)

		FindVideoTotal(ctx context.Context, categoryPid, categoryChildId int64) (count int64, err error)

		//FindVideoListByHot 查询类型对应分类的热播视频
		FindVideoListByHot(ctx context.Context, categoryPid, tabType int64) (data []*Videos, err error)
	}

	defaultVideosModel struct {
		sqlc.CachedConn
		table string
	}

	Videos struct {
		TypeName              sql.NullString `db:"type_ame"`  // 类型名称
		TypeSort              int64          `db:"type_sort"` //类型排序
		Id                    int64          `db:"id"`
		Title                 string         `db:"title"`             // 影片标题
		CategoryPid           int64          `db:"category_pid"`      // 分类一级id
		CategoryChildId       int64          `db:"category_child_id"` // 分类二级id
		SurfacePlot           sql.NullString `db:"surface_plot"`      // 影片封面图
		Recommend             int64          `db:"recommend"`         // 是否推荐 1是 2否
		Cycle                 int64          `db:"cycle"`             // 是否轮播 1是 2否
		CycleImg              sql.NullString `db:"cycle_img"`         // 轮播图片
		ChargingMode          int64          `db:"charging_mode"`     // 收费模式 1免费 2vip免费 3金币点播
		BuyMode               int64          `db:"buy_mode"`          // 购买模式 1按部 2按集
		Gold                  int64          `db:"gold"`              // 金币点播值
		Directors             sql.NullString `db:"directors"`         // 导演
		Actors                sql.NullString `db:"actors"`            // 演员
		ImdbScore             int64          `db:"imdb_score"`        // imd评分.百分制
		ImdbScoreId           string         `db:"imdb_score_id"`     // imd评分ID
		DoubanScore           int64          `db:"douban_score"`      // 豆瓣评分.百分制
		DoubanScoreId         string         `db:"douban_score_id"`   // 豆瓣评分ID
		Introduce             sql.NullString `db:"introduce"`         // 简介
		PopularityDay         int64          `db:"popularity_day"`    // 日人气
		PopularityWeek        int64          `db:"popularity_week"`   // 周人气
		PopularityMonth       int64          `db:"popularity_month"`  // 月人气
		PopularitySum         int64          `db:"popularity_sum"`    // 总人气
		Note                  string         `db:"note"`              // 连载状态
		Year                  string         `db:"year"`              // 年份
		AlbumId               int64          `db:"album_id"`          // 关联专题id
		Status                int64          `db:"status"`            // 状态
		CreateAt              int64          `db:"create_at"`
		UpdateAt              int64          `db:"update_at"`
		Duration              int64          `db:"duration"`          // 时长(单位s)
		Region                string         `db:"region"`            // 自定义地区
		Language              string         `db:"language"`          // 自定义语言
		Label                 string         `db:"label"`             // 自定义标签
		Number                sql.NullInt64  `db:"number"`            // 总集数
		Total                 sql.NullInt64  `db:"total"`             // 更新集数
		HorizontalPoster      sql.NullString `db:"horizontal_poster"` // 横屏海报
		VerticalPoster        sql.NullString `db:"vertical_poster"`   // 竖屏海报
		Publish               sql.NullString `db:"publish"`           // 发行商
		SerialNumber          sql.NullString `db:"serial_number"`     // 序列号
		Screenshot            sql.NullString `db:"screenshot"`        // 截屏
		Gif                   sql.NullString `db:"gif"`
		Alias                 sql.NullString `db:"alias"`
		ReleaseAt             sql.NullInt64  `db:"release_at"`
		ShelfAt               int64          `db:"shelf_at"`
		End                   sql.NullInt64  `db:"end"`
		Unit                  sql.NullString `db:"unit"`
		Watch                 sql.NullInt64  `db:"watch"`
		CollectionId          sql.NullInt64  `db:"collection_id"`
		UseLocalImage         sql.NullInt64  `db:"use_local_image"`
		TitlesTime            int64          `db:"titles_time"`              // 片头时间
		TrailerTime           int64          `db:"trailer_time"`             // 片尾时间
		SiteId                int64          `db:"v.site_id"`                // 站点id
		CategoryPidStatus     int64          `db:"category_pid_status"`      // 顶级分类状态
		CategoryChildIdStatus int64          `db:"category_child_id_status"` // 子级分类状态
		PlayUrl               sql.NullString `db:"play_url"`                 // 采集的源地址
		PlayUrlPutIn          int64          `db:"play_url_put_in"`          // 播放地址是否入库1-已经入库
	}
)

func newVideosModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultVideosModel {
	return &defaultVideosModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`videos`",
	}
}

func (m *defaultVideosModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	cmsVideosIdKey := fmt.Sprintf("%s%v", cacheCmsVideosIdPrefix, id)
	cmsVideosSiteIdCategoryPidCategoryChildIdTitleKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheCmsVideosSiteIdCategoryPidCategoryChildIdTitlePrefix, data.SiteId, data.CategoryPid, data.CategoryChildId, data.Title)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, cmsVideosIdKey, cmsVideosSiteIdCategoryPidCategoryChildIdTitleKey)
	return err
}

func (m *defaultVideosModel) FindOne(ctx context.Context, id int64) (*Videos, error) {
	cmsVideosIdKey := fmt.Sprintf("%s%v", cacheCmsVideosIdPrefix, id)
	var resp Videos
	err := m.QueryRowCtx(ctx, &resp, cmsVideosIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videosRows, m.table)
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

func (m *defaultVideosModel) FindOneBySiteIdCategoryPidCategoryChildIdTitle(ctx context.Context, siteId int64, categoryPid int64, categoryChildId int64, title string) (*Videos, error) {
	cmsVideosSiteIdCategoryPidCategoryChildIdTitleKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheCmsVideosSiteIdCategoryPidCategoryChildIdTitlePrefix, siteId, categoryPid, categoryChildId, title)
	var resp Videos
	err := m.QueryRowIndexCtx(ctx, &resp, cmsVideosSiteIdCategoryPidCategoryChildIdTitleKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `site_id` = ? and `category_pid` = ? and `category_child_id` = ? and `title` = ? limit 1", videosRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, siteId, categoryPid, categoryChildId, title); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, models.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideosModel) Insert(ctx context.Context, data *Videos) (sql.Result, error) {
	cmsVideosIdKey := fmt.Sprintf("%s%v", cacheCmsVideosIdPrefix, data.Id)
	cmsVideosSiteIdCategoryPidCategoryChildIdTitleKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheCmsVideosSiteIdCategoryPidCategoryChildIdTitlePrefix, data.SiteId, data.CategoryPid, data.CategoryChildId, data.Title)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, videosRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Title, data.CategoryPid, data.CategoryChildId, data.SurfacePlot, data.Recommend, data.Cycle, data.CycleImg, data.ChargingMode, data.BuyMode, data.Gold, data.Directors, data.Actors, data.ImdbScore, data.ImdbScoreId, data.DoubanScore, data.DoubanScoreId, data.Introduce, data.PopularityDay, data.PopularityWeek, data.PopularityMonth, data.PopularitySum, data.Note, data.Year, data.AlbumId, data.Status, data.Duration, data.Region, data.Language, data.Label, data.Number, data.Total, data.HorizontalPoster, data.VerticalPoster, data.Publish, data.SerialNumber, data.Screenshot, data.Gif, data.Alias, data.ReleaseAt, data.ShelfAt, data.End, data.Unit, data.Watch, data.CollectionId, data.UseLocalImage, data.TitlesTime, data.TrailerTime, data.SiteId, data.CategoryPidStatus, data.CategoryChildIdStatus, data.PlayUrl, data.PlayUrlPutIn)
	}, cmsVideosIdKey, cmsVideosSiteIdCategoryPidCategoryChildIdTitleKey)
	return ret, err
}

func (m *defaultVideosModel) Update(ctx context.Context, newData *Videos) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	cmsVideosIdKey := fmt.Sprintf("%s%v", cacheCmsVideosIdPrefix, data.Id)
	cmsVideosSiteIdCategoryPidCategoryChildIdTitleKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheCmsVideosSiteIdCategoryPidCategoryChildIdTitlePrefix, data.SiteId, data.CategoryPid, data.CategoryChildId, data.Title)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, videosRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Title, newData.CategoryPid, newData.CategoryChildId, newData.SurfacePlot, newData.Recommend, newData.Cycle, newData.CycleImg, newData.ChargingMode, newData.BuyMode, newData.Gold, newData.Directors, newData.Actors, newData.ImdbScore, newData.ImdbScoreId, newData.DoubanScore, newData.DoubanScoreId, newData.Introduce, newData.PopularityDay, newData.PopularityWeek, newData.PopularityMonth, newData.PopularitySum, newData.Note, newData.Year, newData.AlbumId, newData.Status, newData.Duration, newData.Region, newData.Language, newData.Label, newData.Number, newData.Total, newData.HorizontalPoster, newData.VerticalPoster, newData.Publish, newData.SerialNumber, newData.Screenshot, newData.Gif, newData.Alias, newData.ReleaseAt, newData.ShelfAt, newData.End, newData.Unit, newData.Watch, newData.CollectionId, newData.UseLocalImage, newData.TitlesTime, newData.TrailerTime, newData.SiteId, newData.CategoryPidStatus, newData.CategoryChildIdStatus, newData.PlayUrl, newData.PlayUrlPutIn, newData.Id)
	}, cmsVideosIdKey, cmsVideosSiteIdCategoryPidCategoryChildIdTitleKey)
	return err
}

func (m *defaultVideosModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheCmsVideosIdPrefix, primary)
}

func (m *defaultVideosModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videosRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultVideosModel) tableName() string {
	return m.table
}

// FindRecommendList 查询推荐页面的数据
func (m *defaultVideosModel) FindRecommendList(ctx context.Context) (data []*Videos, err error) {
	//query := fmt.Sprintf("select * from %s where  ", m.table)
	query := `select c.name as type_ame,c.sort as type_sort, v.id,v.title,category_pid,category_child_id,surface_plot,recommend,cycle,cycle_img,charging_mode,buy_mode,gold,directors,actors,
imdb_score,imdb_score_id,douban_score,douban_score_id,introduce,popularity_day,popularity_week,popularity_month,popularity_sum,v.note,year,album_id,v.status,v.create_at,
v.update_at,duration,region,v.language,label,v.number,v.total,horizontal_poster,vertical_poster,publish,serial_number,screenshot,gif,
alias,release_at,shelf_at,end,unit,watch,collection_id,use_local_image,titles_time,trailer_time,v.site_id,category_pid_status,category_child_id_status,play_url,play_url_put_in
from cms.videos as v Left join cms.category as c on v.category_pid=c.id where v.recommend = 1 and c.status=1;`
	var resp []*Videos
	allCmsVideosIdKey := "cache:cms:videos:recommend"
	err = m.GetCacheCtx(ctx, allCmsVideosIdKey, &resp)
	if err != nil || resp == nil {
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query)
		if err == nil {
			err1 := m.SetCacheWithExpireCtx(ctx, allCmsVideosIdKey, resp, time.Hour*12)
			if err1 != nil {
				logx.Error("设置缓存失败了:", err1)
			}
		}
	}
	switch err {
	case nil:
		return resp, err
	case sqlc.ErrNotFound:
		return nil, models.ErrNotFound
	default:
		return nil, err
	}
}

// FindVideoList 查询推荐页面的数据
func (m *defaultVideosModel) FindVideoList(ctx context.Context, pageIndex, pageSize, categoryPid, categoryChildId int64) (data []*Videos, err error) {
	limit := (pageIndex - 1) * pageSize

	query := `select c.name as type_ame,c.sort as type_sort, v.id,v.title,category_pid,category_child_id,surface_plot,recommend,cycle,cycle_img,charging_mode,buy_mode,gold,directors,actors,
imdb_score,imdb_score_id,douban_score,douban_score_id,introduce,popularity_day,popularity_week,popularity_month,popularity_sum,v.note,year,album_id,v.status,v.create_at,
v.update_at,duration,region,v.language,label,v.number,v.total,horizontal_poster,vertical_poster,publish,serial_number,screenshot,gif,
alias,release_at,shelf_at,end,unit,watch,collection_id,use_local_image,titles_time,trailer_time,v.site_id,category_pid_status,category_child_id_status,play_url,play_url_put_in
from cms.videos as v Left join cms.category as c on v.category_pid=c.id 
                          where c.status=1`
	args := make([]any, 0)
	allCmsVideosIdKey := "cache:cms:videos"
	if categoryPid != 0 {
		query += ` and category_pid=?`
		args = append(args, categoryPid)
		allCmsVideosIdKey = fmt.Sprintf("%s:%v", allCmsVideosIdKey, categoryPid)
	}
	if categoryChildId != 0 {
		query += `  and category_child_id=?`
		args = append(args, categoryChildId)
		allCmsVideosIdKey = fmt.Sprintf("%s:%v", allCmsVideosIdKey, categoryChildId)
	}
	query += ` limit ?,? ;`
	args = append(args, limit)
	args = append(args, pageSize)
	var resp []*Videos
	err = m.GetCacheCtx(ctx, allCmsVideosIdKey, &resp)
	if err != nil || resp == nil {
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, args...)
		if err == nil {
			err1 := m.SetCacheWithExpireCtx(ctx, allCmsVideosIdKey, resp, time.Second*30)
			if err1 != nil {
				logx.Error("设置缓存失败了:", err1)
			}
		}
	}
	switch err {
	case nil:
		return resp, err
	case sqlc.ErrNotFound:
		return nil, models.ErrNotFound
	default:
		return nil, err
	}
}

// FindBannerList 查询banner 信息
func (m *defaultVideosModel) FindBannerList(ctx context.Context, categoryPid int64) (data []*Videos, err error) {

	query := `select c.name as type_ame,c.sort as type_sort, v.id,v.title,category_pid,category_child_id,surface_plot,recommend,cycle,cycle_img,charging_mode,buy_mode,gold,directors,actors,
imdb_score,imdb_score_id,douban_score,douban_score_id,introduce,popularity_day,popularity_week,popularity_month,popularity_sum,v.note,year,album_id,v.status,v.create_at,
v.update_at,duration,region,v.language,label,v.number,v.total,horizontal_poster,vertical_poster,publish,serial_number,screenshot,gif,
alias,release_at,shelf_at,end,unit,watch,collection_id,use_local_image,titles_time,trailer_time,v.site_id,category_pid_status,category_child_id_status,play_url,play_url_put_in
from cms.videos as v Left join cms.category as c on v.category_pid=c.id 
                          where c.status=1 and v.cycle = 1`
	args := make([]any, 0)
	allCmsVideosIdKey := "cache:cms:videos:banner"
	if categoryPid != 0 {
		query += ` and category_pid=?`
		args = append(args, categoryPid)
		allCmsVideosIdKey = fmt.Sprintf("%s:%v", allCmsVideosIdKey, categoryPid)
	}
	var resp []*Videos
	err = m.GetCacheCtx(ctx, allCmsVideosIdKey, &resp)
	if err != nil || resp == nil {
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, args...)
		if err == nil {
			err1 := m.SetCacheWithExpireCtx(ctx, allCmsVideosIdKey, resp, time.Second*30)
			if err1 != nil {
				logx.Error("设置缓存失败了:", err1)
			}
		}
	}
	switch err {
	case nil:
		return resp, err
	case sqlc.ErrNotFound:
		return nil, models.ErrNotFound
	default:
		return nil, err
	}
}

// FindVideoTotal 获取总条数
func (m *defaultVideosModel) FindVideoTotal(ctx context.Context, categoryPid, categoryChildId int64) (data int64, err error) {
	query := `select count(*) total from cms.videos as v Left join cms.category as c on v.category_pid=c.id  where c.status=1`
	args := make([]any, 0)
	allCmsVideosIdKey := "cache:cms:videos:total"
	if categoryPid != 0 {
		query += ` and category_pid=?`
		args = append(args, categoryPid)
		allCmsVideosIdKey = fmt.Sprintf("%s:%v", allCmsVideosIdKey, categoryPid)
	}
	if categoryChildId != 0 {
		query += `  and category_child_id=?`
		args = append(args, categoryChildId)
		allCmsVideosIdKey = fmt.Sprintf("%s:%v", allCmsVideosIdKey, categoryChildId)
	}
	var resp VideoTotal
	//err = m.GetCacheCtx(ctx, allCmsVideosIdKey, resp)
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, args...)
	if err == nil {
		err1 := m.SetCacheWithExpireCtx(ctx, allCmsVideosIdKey, resp, time.Second*30)
		if err1 != nil {
			logx.Error("设置缓存失败了:", err1)
		}
	}
	switch err {
	case nil:
		return resp.Total, err
	case sqlc.ErrNotFound:
		return 0, models.ErrNotFound
	default:
		return 0, err
	}
}

// FindVideoListByHot 查询类型对应分类的热播视频
func (m *defaultVideosModel) FindVideoListByHot(ctx context.Context, categoryPid, tabType int64) (data []*Videos, err error) {
	query := `select c.name as type_ame,c.sort as type_sort, v.id,v.title,category_pid,category_child_id,surface_plot,recommend,cycle,cycle_img,charging_mode,buy_mode,gold,directors,actors,
imdb_score,imdb_score_id,douban_score,douban_score_id,introduce,popularity_day,popularity_week,popularity_month,popularity_sum,v.note,year,album_id,v.status,v.create_at,
v.update_at,duration,region,v.language,label,v.number,v.total,horizontal_poster,vertical_poster,publish,serial_number,screenshot,gif,
alias,release_at,shelf_at,end,unit,watch,collection_id,use_local_image,titles_time,trailer_time,v.site_id,category_pid_status,category_child_id_status,play_url,play_url_put_in
from cms.videos as v Left join cms.category as c on v.category_pid=c.id 
                          where c.status=1 and category_pid=? `
	args := make([]any, 0)
	allCmsVideosIdKey := "cache:cms:videos:hot"
	allCmsVideosIdKey = fmt.Sprintf("%s:%v", allCmsVideosIdKey, categoryPid)
	allCmsVideosIdKey = fmt.Sprintf("%s:%v", allCmsVideosIdKey, tabType)
	//query += ` and category_pid=?`
	//args = append(args, categoryPid)

	if tabType == 0 { //天热播
		query += ` order by v.year desc,v.popularity_day desc`
	} else if tabType == 1 { //周热播
		query += ` order by v.year desc,v.popularity_week desc`
	} else if tabType == 2 { //月热播
		query += ` order by v.year desc,v.popularity_month desc`
	} else { //年热播
		query += ` order by v.year desc,v.popularity_sum desc`
	}
	args = append(args, categoryPid)
	query += ` limit ?,?`
	args = append(args, 0, 50)

	logx.Infov(query)
	var resp []*Videos
	err = m.GetCacheCtx(ctx, allCmsVideosIdKey, &resp)
	if err != nil || resp == nil {
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, args...)
		if err == nil {
			err1 := m.SetCacheWithExpireCtx(ctx, allCmsVideosIdKey, resp, time.Minute*10)
			if err1 != nil {
				logx.Error("设置缓存失败了:", err1)
			}
		}
	}
	switch err {
	case nil:
		return resp, err
	case sqlc.ErrNotFound:
		return nil, models.ErrNotFound
	default:
		return nil, err
	}
}
