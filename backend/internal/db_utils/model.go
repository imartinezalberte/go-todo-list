package db_utils

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/imartinezalberte/go-todo-list/backend/server/entities"
)

type Model struct {
	ID        uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4();"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m Model) Res() entities.Model {
	return entities.Model{
		ID:        m.ID,
		CreatedAt: entities.DateYearMonthDayHourMinute(m.CreatedAt),
	}
}
