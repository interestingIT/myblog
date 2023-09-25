// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tag is the golang structure of table tag for DAO operations like Where/Data.
type Tag struct {
	g.Meta     `orm:"table:tag, do:true"`
	Id         interface{} // 标签id
	Name       interface{} // 标签名称
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 最后一次更新时间
	IsDeleted  interface{} // 删除标志位：0：未删除 1：已删除
}
