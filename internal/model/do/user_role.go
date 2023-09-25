// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRole is the golang structure of table user_role for DAO operations like Where/Data.
type UserRole struct {
	g.Meta     `orm:"table:user_role, do:true"`
	Id         interface{} // id
	Username   interface{} // 用户名
	Role       interface{} // 角色
	CreateTime *gtime.Time // 创建时间
}
