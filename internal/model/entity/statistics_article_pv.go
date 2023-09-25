// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// StatisticsArticlePv is the golang structure for table statistics_article_pv.
type StatisticsArticlePv struct {
	Id         uint64      `json:"id"         description:"id"`
	PvDate     *gtime.Time `json:"pvDate"     description:"被统计的日期"`
	PvCount    uint64      `json:"pvCount"    description:"pv浏览量"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"最后一次更新时间"`
}
