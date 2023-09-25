// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleCategoryRelDao is the data access object for table article_category_rel.
type ArticleCategoryRelDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ArticleCategoryRelColumns // columns contains all the column names of Table for convenient usage.
}

// ArticleCategoryRelColumns defines and stores column names for table article_category_rel.
type ArticleCategoryRelColumns struct {
	Id         string // id
	ArticleId  string // 文章id
	CategoryId string // 分类id
}

// articleCategoryRelColumns holds the columns for table article_category_rel.
var articleCategoryRelColumns = ArticleCategoryRelColumns{
	Id:         "id",
	ArticleId:  "article_id",
	CategoryId: "category_id",
}

// NewArticleCategoryRelDao creates and returns a new DAO object for table data access.
func NewArticleCategoryRelDao() *ArticleCategoryRelDao {
	return &ArticleCategoryRelDao{
		group:   "default",
		table:   "article_category_rel",
		columns: articleCategoryRelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ArticleCategoryRelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ArticleCategoryRelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ArticleCategoryRelDao) Columns() ArticleCategoryRelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ArticleCategoryRelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ArticleCategoryRelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ArticleCategoryRelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
