package tag

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/imartinezalberte/go-todo-list/backend/internal/db_utils"
	"github.com/imartinezalberte/go-todo-list/backend/internal/pagination"
	"github.com/imartinezalberte/go-todo-list/backend/server"
)

type (
	TagQuery struct{ ID uuid.UUID }

	TagsQuery struct {
		Tag
		pagination.Query
	}

	AddTagCommand    struct{ Tag }
	UpdateTagCommand struct{ Tag }
	DelTagCommand    struct{ ID uuid.UUID }
)

func (t TagQuery) Scopes() server.ScopeFn {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func (t TagsQuery) Scopes() server.ScopeFn {
	scopes := []func(*gorm.DB) *gorm.DB{t.Query.Scopes()}

	if sanitized := strings.TrimSpace(t.Tag.Name); sanitized != "" {
		scopes = append(scopes, db_utils.Where("name like ?", "%"+sanitized+"%"))
	}

	if sanitized := strings.TrimSpace(t.Tag.Description); sanitized != "" {
		scopes = append(scopes, db_utils.Where("description like ?", "%"+sanitized+"%"))
	}

	if sanitized := strings.TrimSpace(t.Tag.Color); sanitized != "" {
		scopes = append(scopes, db_utils.Where("color like ?", "%"+sanitized+"%"))
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(scopes...)
	}
}

func (t AddTagCommand) Scopes() server.ScopeFn {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func (t UpdateTagCommand) Scopes() server.ScopeFn {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func (t DelTagCommand) Scopes() server.ScopeFn {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
