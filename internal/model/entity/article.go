// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Article is the golang structure for table article.
type Article struct {
	Id          uint64      `json:"id"          description:"文章id"`
	Title       string      `json:"title"       description:"文章标题"`
	TitleImage  string      `json:"titleImage"  description:"文章题图"`
	Description string      `json:"description" description:"文章描述"`
	CreateTime  *gtime.Time `json:"createTime"  description:"创建时间"`
	UpdateTime  *gtime.Time `json:"updateTime"  description:"最后一次更新时间"`
	IsDeleted   int         `json:"isDeleted"   description:"删除标志位：0：未删除 1：已删除"`
	ReadNum     uint        `json:"readNum"     description:"被阅读次数"`
}
