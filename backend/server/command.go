package server

import "gorm.io/gorm"

type (
	Commander interface {
		ToCmd() (Scoper, error)
	}

	ScopeFn func(*gorm.DB) *gorm.DB
	Scoper  interface {
		Scopes() ScopeFn
	}
)
