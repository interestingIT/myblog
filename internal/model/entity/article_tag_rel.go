// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ArticleTagRel is the golang structure for table article_tag_rel.
type ArticleTagRel struct {
	Id        uint64 `json:"id"        description:"id"`
	ArticleId uint64 `json:"articleId" description:"文章id"`
	TagId     uint64 `json:"tagId"     description:"标签id"`
}
