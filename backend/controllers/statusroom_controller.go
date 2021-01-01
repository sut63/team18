package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/statusroom"
)

// StatusRoomController defines the struct for the statusroom controller
type StatusRoomController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateStatusRoom handles POST requests for adding statusroom entities
// @Summary Create statusroom
// @Description Create statusroom
// @ID create-statusroom
// @Accept   json
// @Produce  json
// @Param statusroom body ent.StatusRoom true "StatusRoom entity"
// @Success 200 {object} ent.StatusRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusrooms [post]
func (ctl *StatusRoomController) CreateStatusRoom(c *gin.Context) {
	obj := ent.StatusRoom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "statusroom binding failed",
		})
		return
	}

	s, err := ctl.client.StatusRoom.
		Create().
		SetStatusName(obj.StatusName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, s)
}

// GetStatusRoom handles GET requests to retrieve a statusroom entity
// @Summary Get a statusroom entity by ID
// @Description get statusroom by ID
// @ID get-statusroom
// @Produce  json
// @Param id path int true "StatusRoom ID"
// @Success 200 {object} ent.StatusRoom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusrooms/{id} [get]
func (ctl *StatusRoomController) GetStatusRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	s, err := ctl.client.StatusRoom.
		Query().
		Where(statusroom.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListStatusRoom handles request to get a list of statusroom entities
// @Summary List statusroom entities
// @Description list statusroom entities
// @ID list-statusroom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.StatusRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusrooms [get]
func (ctl *StatusRoomController) ListStatusRoom(c *gin.Context) {
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

	statuss, err := ctl.client.StatusRoom.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, statuss)
}

// NewStatusRoomController creates and registers handles for the statusroom controller
func NewStatusRoomController(router gin.IRouter, client *ent.Client) *StatusRoomController {
	sc := &StatusRoomController{
		client: client,
		router: router,
	}
	sc.register()
	return sc
}

// InitStatusRoomController registers routes to the main engine
func (ctl *StatusRoomController) register() {
	statusrooms := ctl.router.Group("/statusrooms")
	statusrooms.GET("", ctl.ListStatusRoom)
	// CRUD
	statusrooms.POST("", ctl.CreateStatusRoom)
	statusrooms.GET(":id", ctl.GetStatusRoom)
}
