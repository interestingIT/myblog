// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CategoryDao is the data access object for table category.
type CategoryDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns CategoryColumns // columns contains all the column names of Table for convenient usage.
}

// CategoryColumns defines and stores column names for table category.
type CategoryColumns struct {
	Id         string // 标签id
	Name       string // 分类名称
	CreateTime string // 创建时间
	UpdateTime string // 最后一次更新时间
	IsDeleted  string // 删除标志位：0：未删除 1：已删除
}

// categoryColumns holds the columns for table category.
var categoryColumns = CategoryColumns{
	Id:         "id",
	Name:       "name",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	IsDeleted:  "is_deleted",
}

// NewCategoryDao creates and returns a new DAO object for table data access.
func NewCategoryDao() *CategoryDao {
	return &CategoryDao{
		group:   "default",
		table:   "category",
		columns: categoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CategoryDao) Columns() CategoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
