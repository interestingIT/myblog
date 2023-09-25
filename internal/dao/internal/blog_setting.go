// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BlogSettingDao is the data access object for table blog_setting.
type BlogSettingDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns BlogSettingColumns // columns contains all the column names of Table for convenient usage.
}

// BlogSettingColumns defines and stores column names for table blog_setting.
type BlogSettingColumns struct {
	Id           string // id
	BlogName     string // 博客名称
	Author       string // 作者名
	Introduction string // 介绍语
	Avatar       string // 作者头像
	GithubHome   string // GitHub 主页访问地址
	CsdnHome     string // CSDN 主页访问地址
	GiteeHome    string // Gitee 主页访问地址
	ZhihuHome    string // 知乎主页访问地址
}

// blogSettingColumns holds the columns for table blog_setting.
var blogSettingColumns = BlogSettingColumns{
	Id:           "id",
	BlogName:     "blog_name",
	Author:       "author",
	Introduction: "introduction",
	Avatar:       "avatar",
	GithubHome:   "github_home",
	CsdnHome:     "csdn_home",
	GiteeHome:    "gitee_home",
	ZhihuHome:    "zhihu_home",
}

// NewBlogSettingDao creates and returns a new DAO object for table data access.
func NewBlogSettingDao() *BlogSettingDao {
	return &BlogSettingDao{
		group:   "default",
		table:   "blog_setting",
		columns: blogSettingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BlogSettingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BlogSettingDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BlogSettingDao) Columns() BlogSettingColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BlogSettingDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BlogSettingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BlogSettingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
