package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/statuscheckin"
)

// StatusCheckinController defines the struct for the statusReserve controller
type StatusCheckinController struct {
	client *ent.Client
	router gin.IRouter
}

type StatusCheckin struct {
	statusname string
}

// CreateStatusCheckin handles POST requests for adding statuscheckin entities
// @Summary Create statuscheckin
// @Description Create statuscheckin
// @ID create-statuscheckin
// @Accept   json
// @Produce  json
// @Param statuscheckin body ent.Status true "StatusCheckin entity"
// @Success 200 {object} ent.StatusCheckIn
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuscheckins [post]
func (ctl *StatusCheckinController) CreateStatusCheckin(c *gin.Context) {
	obj := StatusCheckin{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "checkin Status binding failed",
		})
		return
	}

	u, err := ctl.client.StatusCheckIn.
		Create().
		SetStatusName(obj.statusname).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// ListStatusCheckin handles request to get a list of statuscheckin entities
// @Summary List statuscheckin entities
// @Description list statuscheckin entities
// @ID list-statuscheckin
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.StatusCheckIn
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuscheckins [get]
func (ctl *StatusCheckinController) ListStatusCheckin(c *gin.Context) {
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

	statuss, err := ctl.client.StatusCheckIn.
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

// GetStatusCheckin handles GET requests to retrieve a statuscheckin entity
// @Summary Get a statuscheckin entity by ID
// @Description get statuscheckin by ID
// @ID get-statuscheckin
// @Produce  json
// @Param id path int true "statuscheckin ID"
// @Success 200 {object} ent.StatusCheckIn
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuscheckins/{id} [get]
func (ctl *StatusCheckinController) GetStatusCheckin(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.StatusCheckIn.
		Query().
		Where(statuscheckin.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// DeleteStatusCheckin handles DELETE requests to delete a statuscheckin entity
// @Summary Delete a statuscheckin entity by ID
// @Description get statuscheckin by ID
// @ID delete-statuscheckin
// @Produce  json
// @Param id path int true "statuscheckin ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuscheckins/{id} [delete]
func (ctl *StatusCheckinController) DeleteStatusCheckin(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.StatusCheckIn.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// NewStatusCheckinController creates and registers handles for the statuscheckin controller
func NewStatusCheckinController(router gin.IRouter, client *ent.Client) *StatusCheckinController {
	sr := &StatusCheckinController{
		client: client,
		router: router,
	}
	sr.register()

	return sr
}

// InitStatusReserveController registers routes to the main engine
func (ctl *StatusCheckinController) register() {
	statuscheckins := ctl.router.Group("/statuscheckins")

	statuscheckins.GET("", ctl.ListStatusCheckin)

	// CRUD
	statuscheckins.POST("", ctl.CreateStatusCheckin)
	statuscheckins.GET(":id", ctl.GetStatusCheckin)
	statuscheckins.DELETE(":id", ctl.DeleteStatusCheckin)
}
