package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/customer"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/fixroom"
	"github.com/team18/app/ent/furnituredetail"
)

// FixRoomController defines the struct for the fixroom controller
type FixRoomController struct {
	client *ent.Client
	router gin.IRouter
}

// FixRoom struct
type FixRoom struct {
	FixDetail       string
	Customer        int
	FurnitureDetail int
	Room            int
}

// CreateFixRoom handles POST requests for adding fixroom entities
// @Summary Create fixroom
// @Description Create fixroom
// @ID create-fixroom
// @Accept   json
// @Produce  json
// @Param fixroom body ent.FixRoom true "FixRoom entity"
// @Success 200 {object} ent.FixRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixrooms [post]
func (ctl *FixRoomController) CreateFixRoom(c *gin.Context) {
	obj := FixRoom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "fixroom binding failed",
		})
		return
	}

	cus, err := ctl.client.Customer.
		Query().
		Where(customer.IDEQ(int(obj.Customer))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "save sec not found",
		})
		return
	}

	fur, err := ctl.client.FurnitureDetail.
		Query().
		Where(furnituredetail.IDEQ(int(obj.FurnitureDetail))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "save subject not found",
		})
		return
	}

	r, err := ctl.client.DataRoom.
		Query().
		Where(dataroom.IDEQ(int(obj.Room))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "save teacher not found",
		})
		return
	}

	s, err := ctl.client.FixRoom.
		Create().
		SetFixDetail(obj.FixDetail).
		SetCustomer(cus).
		SetFurnitureDetail(fur).
		SetRoom(r).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   s,
	})
}

// GetFixRoom handles GET requests to retrieve a fixroom entity
// @Summary Get a fixroom entity by ID
// @Description get fixroom by ID
// @ID get-fixroom
// @Produce  json
// @Param id path int true "FixRoom ID"
// @Success 200 {object} ent.FixRoom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixrooms/{id} [get]
func (ctl *FixRoomController) GetFixRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	s, err := ctl.client.FixRoom.
		Query().
		Where(fixroom.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListFixRoom handles request to get a list of fixroom entities
// @Summary List fixroom entities
// @Description list fixroom entities
// @ID list-fixroom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.FixRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixrooms [get]
func (ctl *FixRoomController) ListFixRoom(c *gin.Context) {
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

	fixrooms, err := ctl.client.FixRoom.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, fixrooms)
}

// DeleteFixRoom handles DELETE requests to delete a fixroom entity
// @Summary Delete a fixroom entity by ID
// @Description get fixroom by ID
// @ID delete-fixroom
// @Produce  json
// @Param id path int true "FixRoom ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /fixrooms/{id} [delete]
func (ctl *FixRoomController) DeleteFixRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.FixRoom.
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

// NewFixRoomController creates and registers handles for the fixroom controller
func NewFixRoomController(router gin.IRouter, client *ent.Client) *FixRoomController {
	pc := &FixRoomController{
		client: client,
		router: router,
	}
	pc.register()
	return pc
}

// InitFixRoomController registers routes to the main engine
func (ctl *FixRoomController) register() {
	fixrooms := ctl.router.Group("/fixrooms")
	fixrooms.GET("", ctl.ListFixRoom)
	// CRUD
	fixrooms.POST("", ctl.CreateFixRoom)
	fixrooms.GET(":id", ctl.GetFixRoom)
	fixrooms.DELETE(":id", ctl.DeleteFixRoom)
}
