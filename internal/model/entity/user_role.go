// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRole is the golang structure for table user_role.
type UserRole struct {
	Id         uint64      `json:"id"         description:"id"`
	Username   string      `json:"username"   description:"用户名"`
	Role       string      `json:"role"       description:"角色"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
}
