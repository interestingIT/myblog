// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleDao is the data access object for table article.
type ArticleDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ArticleColumns // columns contains all the column names of Table for convenient usage.
}

// ArticleColumns defines and stores column names for table article.
type ArticleColumns struct {
	Id          string // 文章id
	Title       string // 文章标题
	TitleImage  string // 文章题图
	Description string // 文章描述
	CreateTime  string // 创建时间
	UpdateTime  string // 最后一次更新时间
	IsDeleted   string // 删除标志位：0：未删除 1：已删除
	ReadNum     string // 被阅读次数
}

// articleColumns holds the columns for table article.
var articleColumns = ArticleColumns{
	Id:          "id",
	Title:       "title",
	TitleImage:  "title_image",
	Description: "description",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	IsDeleted:   "is_deleted",
	ReadNum:     "read_num",
}

// NewArticleDao creates and returns a new DAO object for table data access.
func NewArticleDao() *ArticleDao {
	return &ArticleDao{
		group:   "default",
		table:   "article",
		columns: articleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ArticleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ArticleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ArticleDao) Columns() ArticleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ArticleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ArticleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ArticleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
