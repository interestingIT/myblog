// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"myblog/internal/dao/internal"
)

// internalBlogSettingDao is internal type for wrapping internal DAO implements.
type internalBlogSettingDao = *internal.BlogSettingDao

// blogSettingDao is the data access object for table blog_setting.
// You can define custom methods on it to extend its functionality as you wish.
type blogSettingDao struct {
	internalBlogSettingDao
}

var (
	// BlogSetting is globally public accessible object for table blog_setting operations.
	BlogSetting = blogSettingDao{
		internal.NewBlogSettingDao(),
	}
)

// Fill with you ideas below.
