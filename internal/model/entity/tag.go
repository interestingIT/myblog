// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tag is the golang structure for table tag.
type Tag struct {
	Id         uint64      `json:"id"         description:"标签id"`
	Name       string      `json:"name"       description:"标签名称"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"最后一次更新时间"`
	IsDeleted  int         `json:"isDeleted"  description:"删除标志位：0：未删除 1：已删除"`
}
