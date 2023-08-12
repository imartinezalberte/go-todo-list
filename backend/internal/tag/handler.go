package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/imartinezalberte/go-todo-list/backend/server"
)

const (
	TagsEndpoint = "/tags"
	TagIDParam   = "tag-id"
	TagEndpoint  = TagsEndpoint + "/:" + TagIDParam
)

type handler struct {
	svc Service
}

func Handler(svc Service) server.Handler {
	return &handler{svc}
}

func (h *handler) Routes(c *gin.Engine) {
	c.GET(TagsEndpoint, server.HandleErr(h.getTags))
	c.GET(TagEndpoint, server.HandleErr(h.getTag))
	c.POST(TagsEndpoint, server.HandleErr(h.addTag))
	c.PUT(TagEndpoint, server.HandleErr(h.updateTag))
	c.DELETE(TagEndpoint, server.HandleErr(h.deleteTag))
}

// Tags
//
//	@Title			tags
//	@Summary		Get tags from database
//	@Description	Get tags from database
//	@Tags			tag
//
//	@Param			size		query	int		false	"size of the page that you want to get"			Example(5)	default(10)	minimum(1)	maximum(100)
//	@Param			page		query	int		false	"at which page you want to start counting"	Example(10)	default(1)	minimum(1)
//	@Param			name		query	string	false	"important"																Example(important)
//	@Param			description	query	string	false	"description of the tag (or something that contains)"					Example(food)
//	@Param			color		query	string	false	"Color of the tag, must start with # and has to be hexadecimal value"	Example(#FFFFFF)
//
//	@Produce		json
//	@Success		200	{object}	TagsResponseDTO
//	@Router			/tags [get]
func (h *handler) getTags(c *gin.Context) error {
	var b GetTagsRequestDTO
	if err := c.ShouldBindQuery(&b); err != nil {
		return err
	}

	res, err := Execute(h.svc, c.Request.Context(), b)
	return server.HandleResponse(c, http.StatusOK, res, err)
}

// Get tag
//
//	@Title			Gettag
//	@Summary		Get tag from uuid
//	@Description	Get tag from uuid
//	@Tags			tag
//
//	@Param			tagid	path	string	true	"tagID" Example(4eb4bf58-2c34-4768-8533-40944b4d405a)
//
//	@Produce		json
//	@Success		200	{object}	tagResponseDTO
//	@Router			/tags/:tag-id [get]
func (h *handler) getTag(c *gin.Context) error {
	var b GetTagRequestDTO

	if err := c.ShouldBindUri(&b); err != nil {
		return err
	}

	res, err := Execute(h.svc, c, b)
	return server.HandleResponse(c, http.StatusFound, res, err)
}

// Add tag
//
//	@Title			Addtag
//	@Summary		Add a new tag if it is not already present in the database
//	@Description	Add a new tag if it is not already present in the database
//	@Tags			tag
//
//	@Param			tag	body	tagRequestDTO	true	"tag Request"
//
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	tagResponseDTO
//	@Router			/tags [post]
func (h *handler) addTag(c *gin.Context) error {
	var b AddTagRequestDTO
	if err := c.ShouldBindJSON(&b); err != nil {
		return err
	}

	res, err := Execute(h.svc, c, b)
	return server.HandleResponse(c, http.StatusCreated, res, err)
}

// Update tag
//
//	@Title			Updatetag
//	@Summary		Update tag from uuid
//	@Description	Update tag from uuid
//	@Tags			tag
//
//	@Param			tag-id	path	string			true	"tagID"
//	@Param			tag		body	tagRequestDTO	true	"tagRequest"
//
//	@Produce		json
//	@Success		200	{object}	tagResponseDTO
//	@Router			/tags/:tag-id [put]
func (h *handler) updateTag(c *gin.Context) error {
	var b UpdateTagRequestDTO
	if err := c.ShouldBindUri(&b); err != nil {
		return err
	}

	if err := server.ShouldBindWith(c, &b, binding.JSON); err != nil {
		return err
	}

	res, err := Execute(h.svc, c, b)
	return server.HandleResponse(c, http.StatusOK, res, err)
}

// Delete tag
//
//	@Title			Deletetag
//	@Summary		Delete tag by uuid
//	@Description	Delete tag by uuid
//	@Tags			tag
//
//	@Param			tagid	path	string	true	"tagID"
//
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/tags/:tag-id [delete]
func (h *handler) deleteTag(c *gin.Context) error {
	var b DelTagRequestDTO

	if err := c.ShouldBindUri(&b); err != nil {
		return err
	}

	_, err := Execute(h.svc, c, b)
	return server.HandleResponse(
		c,
		http.StatusAccepted,
		gin.H{"result": "OK"},
		err,
	)
}
