package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/typeroom"
)

// TypeRoomController defines the struct for the typeroom controller
type TypeRoomController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateTypeRoom handles POST requests for adding typeroom entities
// @Summary Create typeroom
// @Description Create typeroom
// @ID create-typeroom
// @Accept   json
// @Produce  json
// @Param typeroom body ent.TypeRoom true "TypeRoom entity"
// @Success 200 {object} ent.TypeRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typerooms [post]
func (ctl *TypeRoomController) CreateTypeRoom(c *gin.Context) {
	obj := ent.TypeRoom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "typeroom binding failed",
		})
		return
	}

	t, err := ctl.client.TypeRoom.
		Create().
		SetTypeName(obj.TypeName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, t)
}

// GetTypeRoom handles GET requests to retrieve a typeroom entity
// @Summary Get a typeroom entity by ID
// @Description get typeroom by ID
// @ID get-typeroom
// @Produce  json
// @Param id path int true "TypeRoom ID"
// @Success 200 {object} ent.TypeRoom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typerooms/{id} [get]
func (ctl *TypeRoomController) GetTypeRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	t, err := ctl.client.TypeRoom.
		Query().
		Where(typeroom.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, t)
}

// ListTypeRoom handles request to get a list of typeroom entities
// @Summary List typeroom entities
// @Description list typeroom entities
// @ID list-typeroom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.TypeRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typerooms [get]
func (ctl *TypeRoomController) ListTypeRoom(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	typerooms, err := ctl.client.TypeRoom.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, typerooms)
}

// NewTypeRoomController creates and registers handles for the typeroom controller
func NewTypeRoomController(router gin.IRouter, client *ent.Client) *TypeRoomController {
	tc := &TypeRoomController{
		client: client,
		router: router,
	}
	tc.register()
	return tc
}

// InitTypeRoomController registers routes to the main engine
func (ctl *TypeRoomController) register() {
	typerooms := ctl.router.Group("/typerooms")
	typerooms.GET("", ctl.ListTypeRoom)
	// CRUD
	typerooms.POST("", ctl.CreateTypeRoom)
	typerooms.GET(":id", ctl.GetTypeRoom)
}
