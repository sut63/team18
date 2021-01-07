package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/dataroom"
)

// DataRoomController defines the struct for the dataroom controller
type DataRoomController struct {
	client *ent.Client
	router gin.IRouter
}
type DataRoom struct {
	RoomNumber string
	Price      float64
	TypeRoom   int
	StatusRoom int
	Promotion  int
}

// CreateDataRoom handles POST requests for adding dataroom entities
// @Summary Create dataroom
// @Description Create dataroom
// @ID create-dataroom
// @Accept   json
// @Produce  json
// @Param dataroom body ent.DataRoom true "DataRoom entity"
// @Success 200 {object} ent.DataRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /datarooms [post]
func (ctl *DataRoomController) CreateDataRoom(c *gin.Context) {
	obj := DataRoom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "dataroom binding failed",
		})
		return
	}

	d, err := ctl.client.DataRoom.
		Create().
		SetPrice(obj.Price).
		SetRoomnumber(obj.RoomNumber).
		SetPromotionID(obj.Promotion).
		SetStatusroomID(obj.StatusRoom).
		SetTyperoomID(obj.TypeRoom).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, d)
}

// GetDataRoom handles GET requests to retrieve a dataroom entity
// @Summary Get a dataroom entity by ID
// @Description get dataroom by ID
// @ID get-dataroom
// @Produce  json
// @Param id path int true "DataRoom ID"
// @Success 200 {object} ent.DataRoom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /datarooms/{id} [get]
func (ctl *DataRoomController) GetDataRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	s, err := ctl.client.DataRoom.
		Query().
		WithPromotion().
		WithStatusroom().
		WithTyperoom().
		Where(dataroom.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListDataRoom handles request to get a list of dataroom entities
// @Summary List dataroom entities
// @Description list dataroom entities
// @ID list-dataroom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.DataRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /datarooms [get]
func (ctl *DataRoomController) ListDataRoom(c *gin.Context) {
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

	datarooms, err := ctl.client.DataRoom.
		Query().
		WithPromotion().
		WithStatusroom().
		WithTyperoom().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, datarooms)
}

// NewDataRoomController creates and registers handles for the dataroomn controller
func NewDataRoomController(router gin.IRouter, client *ent.Client) *DataRoomController {
	dc := &DataRoomController{
		client: client,
		router: router,
	}
	dc.register()
	return dc
}

// InitDataRoomController registers routes to the main engine
func (ctl *DataRoomController) register() {
	datarooms := ctl.router.Group("/datarooms")
	datarooms.GET("", ctl.ListDataRoom)
	// CRUD
	datarooms.POST("", ctl.CreateDataRoom)
	datarooms.GET(":id", ctl.GetDataRoom)
}
