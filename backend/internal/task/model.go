package task

import (
	"github.com/imartinezalberte/go-todo-list/backend/internal/db_utils"
	"github.com/imartinezalberte/go-todo-list/backend/internal/tag"
)

type (
	Task struct {
		db_utils.Model
		Title       string `gorm:"not_null"`
		Description string
		Tags        tag.Tags
	}
)
