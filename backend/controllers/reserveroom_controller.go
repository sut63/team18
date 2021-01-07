<<<<<<< HEAD
package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/customer"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/promotion"
	"github.com/team18/app/ent/reserveroom"
)

// CourseItemController defines the struct for the course item controller
type ReserveRoomController struct {
	client *ent.Client
	router gin.IRouter
}

// Resreve_Room struct
type Resreve_Room struct {
	Rooms       int
	Promotions  int
	Customers   int
	ReserveDate string
	OutDate     string
	NetPrice    float64
}

// CreateReserveRoom handles POST requests for adding ReserveRoom entities
// @Summary Create ReserveRoom
// @Description Create ReserveRoom
// @ID create-ReserveRoom
// @Accept   json
// @Produce  json
// @Param ReserveRoom body ent.ReserveRoom true "ReserveRoom entity"
// @Success 200 {object} ent.ReserveRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ReserveRooms [post]
func (ctl *ReserveRoomController) CreateReserveRoom(c *gin.Context) {
	obj := Resreve_Room{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Resreve Room binding failed",
		})
		return
	}

	cus, err := ctl.client.Customer.
		Query().
		Where(customer.IDEQ(int(obj.Customers))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Customer not found",
		})
		return
	}

	p, err := ctl.client.Promotion.
		Query().
		Where(promotion.IDEQ(int(obj.Promotions))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Promotion not found",
		})
		return
	}

	d, err := ctl.client.DataRoom.
		Query().
		Where(dataroom.IDEQ(int(obj.Rooms))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Room not found",
		})
		return
	}

	timereserve, err := time.Parse(time.RFC3339, obj.ReserveDate)
	timeout, err := time.Parse(time.RFC3339, obj.OutDate)

	r, err := ctl.client.ReserveRoom.
		Create().
		SetReserveDate(timereserve).
		SetDateOut(timeout).
		SetCustomer(cus).
		SetPromotion(p).
		SetRoom(d).
		SetNetPrice(obj.NetPrice).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   r,
	})
}

// GetReserveRoom handles GET requests to retrieve a ReserveRoom entity
// @Summary Get a ReserveRoom entity by ID
// @Description get ReserveRoom by ID
// @ID get-ReserveRoom
// @Produce  json
// @Param id path int true "ReserveRoom ID"
// @Success 200 {object} ent.ReserveRoom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ReserveRooms/{id} [get]
func (ctl *ReserveRoomController) GetReserveRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	i, err := ctl.client.ReserveRoom.
		Query().
		Where(reserveroom.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, i)
}

// ListReserveRoom handles request to get a list of ReserveRoom entities
// @Summary List ReserveRoom entities
// @Description list ReserveRoom entities
// @ID list-ReserveRoom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.ReserveRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ReserveRooms [get]
func (ctl *ReserveRoomController) ListReserveRoom(c *gin.Context) {
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

	reserves, err := ctl.client.ReserveRoom.
		Query().
		WithRoom().
		WithCustomer().
		WithPromotion().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, reserves)
}

// DeleteReserveRoom handles DELETE requests to delete a ReserveRoom entity
// @Summary Delete a ReserveRoom entity by ID
// @Description get ReserveRoom by ID
// @ID delete-ReserveRoom
// @Produce  json
// @Param id path int true "ReserveRooms ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ReserveRooms/{id} [delete]
func (ctl *ReserveRoomController) DeleteReserveRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.ReserveRoom.
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

// UpdateReserveRoom handles PUT requests to update a ReserveRoom entity
// @Summary Update a ReserveRoom entity by ID
// @Description update ReserveRoom by ID
// @ID update-ReserveRoom
// @Accept   json
// @Produce  json
// @Param id path int true "ReserveRoom ID"
// @Param ReserveRoom body ent.ReserveRoom true "ReserveRoom entity"
// @Success 200 {object} ent.ReserveRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ReserveRooms/{id} [put]
func (ctl *ReserveRoomController) UpdateReserveRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.ReserveRoom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Resreve binding failed",
		})
		return
	}
	obj.ID = int(id)
	ci, err := ctl.client.ReserveRoom.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, ci)
}

// NewReserveRoomControllercreates and registers handles for the ReserveRoom controller
func NewReserveRoomController(router gin.IRouter, client *ent.Client) *ReserveRoomController {
	rr := &ReserveRoomController{
		client: client,
		router: router,
	}

	rr.register()

	return rr

}

// InitReserveRoomController registers routes to the main engine
func (ctl *ReserveRoomController) register() {
	ReserveRooms := ctl.router.Group("/ReserveRooms")

	ReserveRooms.GET("", ctl.ListReserveRoom)

	// CRUD
	ReserveRooms.POST("", ctl.CreateReserveRoom)
	ReserveRooms.GET(":id", ctl.GetReserveRoom)
	ReserveRooms.PUT(":id", ctl.UpdateReserveRoom)
	ReserveRooms.DELETE(":id", ctl.DeleteReserveRoom)
}
=======

>>>>>>> 62645723e4c323463db55139e0fe0d2b809a3bd8
