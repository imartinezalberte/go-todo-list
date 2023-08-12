package tag

import "github.com/imartinezalberte/go-todo-list/backend/internal/db_utils"

type (
	Tags []Tag

	Tag struct {
		db_utils.Model
		Name        string `gorm:"unique"`
		Description string `gorm:"not_null"`
		Color       string `gorm:"not_null"`
	}
)

func (t Tag) Map() TagResponseDTO {
	return TagResponseDTO{
		Model:       entities.Model{},
		Name:        t.Name,
		Description: t.Description,
		Color:       t.Color,
	}
}
