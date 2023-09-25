// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StatisticsArticlePvDao is the data access object for table statistics_article_pv.
type StatisticsArticlePvDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns StatisticsArticlePvColumns // columns contains all the column names of Table for convenient usage.
}

// StatisticsArticlePvColumns defines and stores column names for table statistics_article_pv.
type StatisticsArticlePvColumns struct {
	Id         string // id
	PvDate     string // 被统计的日期
	PvCount    string // pv浏览量
	CreateTime string // 创建时间
	UpdateTime string // 最后一次更新时间
}

// statisticsArticlePvColumns holds the columns for table statistics_article_pv.
var statisticsArticlePvColumns = StatisticsArticlePvColumns{
	Id:         "id",
	PvDate:     "pv_date",
	PvCount:    "pv_count",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewStatisticsArticlePvDao creates and returns a new DAO object for table data access.
func NewStatisticsArticlePvDao() *StatisticsArticlePvDao {
	return &StatisticsArticlePvDao{
		group:   "default",
		table:   "statistics_article_pv",
		columns: statisticsArticlePvColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *StatisticsArticlePvDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *StatisticsArticlePvDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *StatisticsArticlePvDao) Columns() StatisticsArticlePvColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *StatisticsArticlePvDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *StatisticsArticlePvDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *StatisticsArticlePvDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
