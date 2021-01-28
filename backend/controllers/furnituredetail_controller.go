package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/counterstaff"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/furniture"
	"github.com/team18/app/ent/furnituredetail"
)

// FurnitureDetailController defines the struct for the furnituredetail controller
type FurnitureDetailController struct {
	client *ent.Client
	router gin.IRouter
}

// FurnitureDetail struct
type FurnitureDetail struct {
	DateAdd         string
	CounterStaff    int
	Furniture       int
	Dataroom        int
	FurnitureAmount int
	FurnitureColour string
	Detail          string
}

// CreateFurnitureDetail handles POST requests for adding furnituredetail entities
// @Summary Create furnituredetail
// @Description Create furnituredetail
// @ID create-furnituredetail
// @Accept   json
// @Produce  json
// @Param furnituredetail body ent.FurnitureDetail true "FurnitureDetail entity"
// @Success 200 {object} ent.FurnitureDetail
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /furnituredetails [post]
func (ctl *FurnitureDetailController) CreateFurnitureDetail(c *gin.Context) {
	obj := FurnitureDetail{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "furnituredetail binding failed",
		})
		return
	}

	settime := time.Now().Format("2006-01-02T15:04:05Z07:00")
	time, err := time.Parse(time.RFC3339, settime)

	cs, err := ctl.client.CounterStaff.
		Query().
		Where(counterstaff.IDEQ(int(obj.CounterStaff))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "counterstaff not found",
		})
		return
	}

	fur, err := ctl.client.Furniture.
		Query().
		Where(furniture.IDEQ(int(obj.Furniture))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Furniture not found",
		})
		return
	}

	dr, err := ctl.client.DataRoom.
		Query().
		Where(dataroom.IDEQ(int(obj.Dataroom))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "DataRoom not found",
		})
	}

	fd, err := ctl.client.FurnitureDetail.
		Create().
		SetDateAdd(time).
		SetCounterstaffs(cs).
		SetFurnitures(fur).
		SetRooms(dr).
		SetFurnitureAmount(obj.FurnitureAmount).
		SetFurnitureColour(obj.FurnitureColour).
		SetDetail(obj.Detail).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error":  err,
			"Status": false,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   fd,
	})
}

// GetFurnitureDetail handles GET requests to retrieve a furnituredetail entity
// @Summary Get a furnituredetail entity by ID
// @Description get furnituredetail by ID
// @ID get-furnituredetail
// @Produce  json
// @Param id path int true "FurnitureDetail ID"
// @Success 200 {object} ent.FurnitureDetail
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /furnituredetails/{id} [get]
func (ctl *FurnitureDetailController) GetFurnitureDetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	s, err := ctl.client.FurnitureDetail.
		Query().
		Where(furnituredetail.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListFurnitureDetail handles request to get a list of furnituredetail entities
// @Summary List furnituredetail entities
// @Description list furnituredetail entities
// @ID list-furnituredetail
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.FurnitureDetail
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /furnituredetails [get]
func (ctl *FurnitureDetailController) ListFurnitureDetail(c *gin.Context) {
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

	furnituredetails, err := ctl.client.FurnitureDetail.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, furnituredetails)
}

// ListFurnitureDetailRoom handles request to get a list of FurnitureDetailRoom entities
// @Summary List FurnitureDetailRoom entities
// @Description list FurnitureDetailRoom entities
// @ID list-FurnitureDetailRoom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Param id path int true "FurnitureDetailRooms ID"
// @Success 200 {array} ent.FurnitureDetail
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /FurnitureDetailRooms/{id} [get]
func (ctl *FurnitureDetailController) ListFurnitureDetailRooms(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	furnituredetails, err := ctl.client.FurnitureDetail.
		Query().
		WithFurnitures().
		WithCounterstaffs().
		WithRooms().
		Limit(limit).
		Offset(offset).
		Where(furnituredetail.HasRoomsWith(dataroom.IDEQ(int(id)))).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, furnituredetails)
}

// NewFurnitureDetailController creates and registers handles for the furnituredetail controller
func NewFurnitureDetailController(router gin.IRouter, client *ent.Client) *FurnitureDetailController {
	pc := &FurnitureDetailController{
		client: client,
		router: router,
	}
	pc.register()
	return pc
}

// InitFurnitureDetailController registers routes to the main engine
func (ctl *FurnitureDetailController) register() {
	furnituredetails := ctl.router.Group("/furnituredetails")
	furnituredetailrooms := ctl.router.Group("/FurnitureDetailRooms")

	furnituredetails.GET("", ctl.ListFurnitureDetail)
	furnituredetailrooms.GET(":id", ctl.ListFurnitureDetailRooms)
	// CRUD
	furnituredetails.POST("", ctl.CreateFurnitureDetail)
	furnituredetails.GET(":id", ctl.GetFurnitureDetail)
}
