// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta     `orm:"table:user, do:true"`
	Id         interface{} // id
	Username   interface{} // 用户名
	Password   interface{} // 密码
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 最后一次更新时间
	IsDeleted  interface{} // 删除标志位：0：未删除 1：已删除
}
