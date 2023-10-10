/*
* author: faith
* date: 2023-9-25
* Description: article show
 */
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ArticlePublishReq struct {
	g.Meta      `path:"/api/admin/article/publish" method:"post"`
	Title       string   `v:"required#标题不能为空"`
	Content     string   `v:"required#内容不能为空"`
	TitleImage  string   `v:"required#标题图片不能为空"`
	Description string   `v:"required#描述不能为空"`
	CategoryId  int      `v:"required|min:1#分类id不能为空|分类id不能为空"`
	Tags        []string `v:"required#标签不能为空"`
}

type ArticlePublishRes struct {
}

type ArticleDelReq struct {
	g.Meta `path:"/api/admin/article/delete/{id}" method:"post"`
}

type ArticleDelRes struct {
}

type ArticleUpdateReq struct {
	g.Meta      `path:"/api/admin/article/update/{id}" method:"post"`
	Id          int      `v:"required|min:1#文章id不能为空|文章id不能为空"`
	Title       string   `v:"required#标题不能为空"`
	Content     string   `v:"required#内容不能为空"`
	TitleImage  string   `v:"required#头图不能为空"`
	CategoryId  int      `v:"required|min:1#分类id不能为空|分类id不能为空"`
	Description string   `v:"required#描述不能为空"`
	Tags        []string `v:"required#标签不能为空"`
}

type ArticleUpdateRes struct {
}

type ArticleListReq struct {
	g.Meta `path:"/api/admin/article/list" method:"post" sm:"search list"`
}

type ArticleListRes struct {
}

type ArticleDetailReq struct {
	g.Meta    `path:"/api/admin/article/detail" method:"post"`
	ArticleId int `v:"required|min:1#文章id不能为空|文章id不能为空"`
}

type ArticleDetailRes struct {
	Id          int
	Title       string
	TitleImage  string
	Content     string
	Description string
	CategoryId  int
	TagIds      []int
}
