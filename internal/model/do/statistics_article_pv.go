// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StatisticsArticlePv is the golang structure of table statistics_article_pv for DAO operations like Where/Data.
type StatisticsArticlePv struct {
	g.Meta     `orm:"table:statistics_article_pv, do:true"`
	Id         interface{} // id
	PvDate     *gtime.Time // 被统计的日期
	PvCount    interface{} // pv浏览量
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 最后一次更新时间
}
