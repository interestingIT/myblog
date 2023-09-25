// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BlogSetting is the golang structure for table blog_setting.
type BlogSetting struct {
	Id           uint64 `json:"id"           description:"id"`
	BlogName     string `json:"blogName"     description:"博客名称"`
	Author       string `json:"author"       description:"作者名"`
	Introduction string `json:"introduction" description:"介绍语"`
	Avatar       string `json:"avatar"       description:"作者头像"`
	GithubHome   string `json:"githubHome"   description:"GitHub 主页访问地址"`
	CsdnHome     string `json:"csdnHome"     description:"CSDN 主页访问地址"`
	GiteeHome    string `json:"giteeHome"    description:"Gitee 主页访问地址"`
	ZhihuHome    string `json:"zhihuHome"    description:"知乎主页访问地址"`
}
