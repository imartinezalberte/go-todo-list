package tag

import (
	"github.com/google/uuid"

	"github.com/imartinezalberte/go-todo-list/backend/internal/db_utils"
	"github.com/imartinezalberte/go-todo-list/backend/internal/pagination"
	"github.com/imartinezalberte/go-todo-list/backend/server"
	"github.com/imartinezalberte/go-todo-list/backend/server/entities"
)

type (
	// Request
	AddTagRequestDTO struct {
		Name        string `json:"name"        binding:"required,min=2"    example:"food"`
		Description string `json:"description" binding:"required"          example:"all the stuff that you love to eat (or not)"`
		Color       string `json:"color"       binding:"required,hexcolor" example:"#FFFFFF"`
	}

	GetTagsRequestDTO struct {
		Name        string `query:"name"        form:"name"        example:"food"`
		Description string `query:"description" form:"description" example:"all the stuff that you love to eat (or not)"`
		Color       string `query:"color"       form:"color"       example:"#FFFFFF"                                     binding:"omitempty,hexcolor"`
		pagination.RequestDTO
	}

	GetTagRequestDTO struct {
		ID string `uri:"tag-id" form:"tag-id" binding:"required,uuid" example:"d8ccd61b-9028-4aa2-8fbb-2cc09be30f6d"`
	}

	UpdateTagRequestDTO struct {
		ID          string `binding:"required,uuid"      example:"d8ccd61b-9028-4aa2-8fbb-2cc09be30f6d"        uri:"tag-id" form:"tag-id"`
		Name        string `binding:"omitempty,min=2"    example:"food"                                                                   json:"name"`
		Description string `                             example:"all the stuff that you love to eat (or not)"                            json:"description"`
		Color       string `binding:"omitempty,hexcolor" example:"#FFFFFF"                                                                json:"color"`
	}

	DelTagRequestDTO struct {
		ID string `uri:"tag-id" form:"tag-id" binding:"required,uuid" example:"d8ccd61b-9028-4aa2-8fbb-2cc09be30f6d"`
	}

	// Response
	TagsResponseDTO []TagResponseDTO

	TagResponseDTO struct {
		entities.Model
		Name        string `json:"name"        example:"food"`
		Description string `json:"description" example:"all the stuff that you love to eat (or not)"`
		Color       string `json:"color"       example:"#FFFFFF"`
	}
)

func (t GetTagsRequestDTO) ToCmd() (server.Scoper, error) {
	return TagsQuery{
		Tag{Name: t.Name, Description: t.Description, Color: t.Color},
		t.RequestDTO.ToCmd(),
	}, nil
}

func (t GetTagRequestDTO) ToCmd() (server.Scoper, error) {
	return TagQuery{uuid.MustParse(t.ID)}, nil
}

func (t AddTagRequestDTO) ToCmd() (server.Scoper, error) {
	return AddTagCommand{Tag{Name: t.Name, Description: t.Description, Color: t.Color}}, nil
}

func (t UpdateTagRequestDTO) ToCmd() (server.Scoper, error) {
	return UpdateTagCommand{
		Tag{
			db_utils.Model{ID: uuid.MustParse(t.ID)},
			t.Name,
			t.Description,
			t.Color,
		},
	}, nil
}

func (t DelTagRequestDTO) ToCmd() (server.Scoper, error) {
	return DelTagCommand{uuid.MustParse(t.ID)}, nil
}
