package db_utils

import "gorm.io/gorm"

func Where(query string, args ...any) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}
