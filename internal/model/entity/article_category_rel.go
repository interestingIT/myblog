// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ArticleCategoryRel is the golang structure for table article_category_rel.
type ArticleCategoryRel struct {
	Id         uint64 `json:"id"         description:"id"`
	ArticleId  uint64 `json:"articleId"  description:"文章id"`
	CategoryId uint64 `json:"categoryId" description:"分类id"`
}
