package pagination

import "gorm.io/gorm"

type Query struct {
	Size int
	Page int
}

func (q Query) Scopes() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(q.Size).Offset((q.Page - 1) * q.Size)
	}
}
