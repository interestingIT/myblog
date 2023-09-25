// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id         uint64      `json:"id"         description:"id"`
	Username   string      `json:"username"   description:"用户名"`
	Password   string      `json:"password"   description:"密码"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"最后一次更新时间"`
	IsDeleted  int         `json:"isDeleted"  description:"删除标志位：0：未删除 1：已删除"`
}
