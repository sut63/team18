package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/statusreserve"
)

// StatusReserveController defines the struct for the statusReserve controller
type StatusReserveController struct {
	client *ent.Client
	router gin.IRouter
}

type StatusReserve struct {
	status string
}

// CreateStatusReserve handles POST requests for adding statusReserve entities
// @Summary Create statusReserve
// @Description Create statusReserve
// @ID create-statusReserve
// @Accept   json
// @Produce  json
// @Param statusReserve body ent.Status true "StatusReserve entity"
// @Success 200 {object} ent.StatusReserve
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusReserves [post]
func (ctl *StatusReserveController) CreateStatusReserve(c *gin.Context) {
	obj := StatusReserve{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Reserve Status binding failed",
		})
		return
	}

	u, err := ctl.client.StatusReserve.
		Create().
		SetStatusName(obj.status).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// ListStatusReserve handles request to get a list of statusReserve entities
// @Summary List statusReserve entities
// @Description list statusReserve entities
// @ID list-statusReserve
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.StatusReserve
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusReserves [get]
func (ctl *StatusReserveController) ListStatusReserve(c *gin.Context) {
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

	statuss, err := ctl.client.StatusReserve.
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

// GetStatusReserve handles GET requests to retrieve a statusReserve entity
// @Summary Get a statusReserve entity by ID
// @Description get statusReserve by ID
// @ID get-statusReserve
// @Produce  json
// @Param id path int true "StatusReserve ID"
// @Success 200 {object} ent.StatusReserve
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusReserves/{id} [get]
func (ctl *StatusReserveController) GetStatusReserve(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.StatusReserve.
		Query().
		Where(statusreserve.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// DeleteStatusReserve handles DELETE requests to delete a statusReserve entity
// @Summary Delete a statusReserve entity by ID
// @Description get statusReserve by ID
// @ID delete-statusReserve
// @Produce  json
// @Param id path int true "StatusReserve ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusReserves/{id} [delete]
func (ctl *StatusReserveController) DeleteStatusReserve(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.StatusReserve.
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

// NewStatusReserveController creates and registers handles for the statusReserve controller
func NewStatusReserveController(router gin.IRouter, client *ent.Client) *StatusReserveController {
	sr := &StatusReserveController{
		client: client,
		router: router,
	}
	sr.register()

	return sr
}

// InitStatusReserveController registers routes to the main engine
func (ctl *StatusReserveController) register() {
	statusReserves := ctl.router.Group("/statusReserves")

	statusReserves.GET("", ctl.ListStatusReserve)

	// CRUD
	statusReserves.POST("", ctl.CreateStatusReserve)
	statusReserves.GET(":id", ctl.GetStatusReserve)
	statusReserves.DELETE(":id", ctl.DeleteStatusReserve)
}
