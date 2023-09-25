// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BlogSetting is the golang structure of table blog_setting for DAO operations like Where/Data.
type BlogSetting struct {
	g.Meta       `orm:"table:blog_setting, do:true"`
	Id           interface{} // id
	BlogName     interface{} // 博客名称
	Author       interface{} // 作者名
	Introduction interface{} // 介绍语
	Avatar       interface{} // 作者头像
	GithubHome   interface{} // GitHub 主页访问地址
	CsdnHome     interface{} // CSDN 主页访问地址
	GiteeHome    interface{} // Gitee 主页访问地址
	ZhihuHome    interface{} // 知乎主页访问地址
}
