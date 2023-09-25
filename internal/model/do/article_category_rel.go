// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleCategoryRel is the golang structure of table article_category_rel for DAO operations like Where/Data.
type ArticleCategoryRel struct {
	g.Meta     `orm:"table:article_category_rel, do:true"`
	Id         interface{} // id
	ArticleId  interface{} // 文章id
	CategoryId interface{} // 分类id
}
