/*
* author: faith
* date: 2023-10-9
* Description: categroy
 */
package v1

import (
	"myblog/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AdminCategoryAddReq struct {
	g.Meta `path:"/api/admin/category/add"  method:"post"`
	Name   string `v:"required#分类名称不能为空"`
}

type AdminCategoryAddRes struct {
}

type AdminCategoryListReq struct {
	g.Meta       `path:"/api/admin/category/list"  method:"post"`
	Current      int `v:"required#当前页不能为空"`
	Size         int `v:"required#每页条数不能为空"`
	StartDate    *gtime.Time
	EndDate      *gtime.Time
	CategoryName string
}

type AdminCategoryListRes struct {
	*entity.Category
}

type AdminCategoryDeleteReq struct {
	g.Meta     `path:"/api/admin/category/delete"  method:"post"`
	CategoryId int `v:"required#分类ID不能为空"`
}

type AdminCategoryDeleteRes struct{}

type AdminCategorySelectListReq struct {
	g.Meta `path:"/api/admin/category/select/list"  method:"post"`
}
