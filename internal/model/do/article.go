// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Article is the golang structure of table article for DAO operations like Where/Data.
type Article struct {
	g.Meta      `orm:"table:article, do:true"`
	Id          interface{} // 文章id
	Title       interface{} // 文章标题
	TitleImage  interface{} // 文章题图
	Description interface{} // 文章描述
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 最后一次更新时间
	IsDeleted   interface{} // 删除标志位：0：未删除 1：已删除
	ReadNum     interface{} // 被阅读次数
}
