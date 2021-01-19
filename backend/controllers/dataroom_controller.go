package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/promotion"
	"github.com/team18/app/ent/reserveroom"
	"github.com/team18/app/ent/statusroom"
)

// DataRoomController defines the struct for the dataroom controller
type DataRoomController struct {
	client *ent.Client
	router gin.IRouter
}

// DataRoom ...
type DataRoom struct {
	RoomNumber string
	Price      float64
	TypeRoom   int
	StatusRoom int
	Promotion  int
	RoomDetail string
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
			"error": "DataRoom binding failed",
		})
		return
	}

	d, err := ctl.client.DataRoom.
		Create().
		SetPrice(obj.Price).
		SetRoomdetail(obj.RoomDetail).
		SetRoomnumber(obj.RoomNumber).
		SetPromotionID(obj.Promotion).
		SetStatusroomID(obj.StatusRoom).
		SetTyperoomID(obj.TypeRoom).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	//success
	c.JSON(200, gin.H{
		"status": false,
		"data":   d,
	})
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

// GetDataroomCustomer handles GET requests to retrieve a dataroomcustomer entity
// @Summary Get a dataroomcustomer entity by ID
// @Description get dataroomcustomer by ID
// @ID get-dataroomcustomer
// @Produce  json
// @Param id path int true "dataroomcustomer ID"
// @Success 200 {object} ent.DataRoom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dataroomcustomer/{id} [get]
func (ctl *DataRoomController) GetDataroomCustomer(c *gin.Context) {
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
		Where(dataroom.HasReservesWith(reserveroom.IDEQ(int(id)))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListDataRoomPromo handles request to get a list of DataRoomPromo entities
// @Summary List DataRoomPromo entities by id
// @Description list DataRoomPromo entities by id
// @ID list-DataRoomPromo
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Param id path int true "DataRoomPromo ID"
// @Success 200 {array} ent.DataRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dataroompromos/{id} [get]
func (ctl *DataRoomController) ListDataRoomPromo(c *gin.Context) {
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

	dataroomspromo, err := ctl.client.DataRoom.
		Query().
		WithPromotion().
		WithStatusroom().
		WithTyperoom().
		Limit(limit).
		Offset(offset).
		Where(dataroom.HasPromotionWith(promotion.IDEQ(int(id))), dataroom.HasStatusroomWith(statusroom.IDEQ(1))).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dataroomspromo)
}

// UpdateDataroom handles PUT requests to update a Dataroom entity
// @Summary Update a Dataroom entity by ID
// @Description update Dataroom by ID
// @ID update-Dataroom
// @Accept   json
// @Produce  json
// @Param id path int true "Dataroom ID"
// @Param Dataroom body ent.DataRoom true "Dataroom entity"
// @Success 200 {object} ent.DataRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /datarooms/{id} [put]
func (ctl *DataRoomController) UpdateDataRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.DataRoom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "room binding failed",
		})
		return

	}
	obj.ID = int(id)

	u, err := ctl.client.DataRoom.
		UpdateOne(&obj).
		SetStatusroom(obj.Edges.Statusroom).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
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
	dataroomcustomer := ctl.router.Group("/dataroomcustomer")
	dataroompromos := ctl.router.Group("/dataroompromos")

	datarooms.GET("", ctl.ListDataRoom)
	dataroompromos.GET(":id", ctl.ListDataRoomPromo)

	// CRUD
	datarooms.POST("", ctl.CreateDataRoom)
	datarooms.GET(":id", ctl.GetDataRoom)
	datarooms.PUT(":id", ctl.UpdateDataRoom)
	dataroomcustomer.GET(":id", ctl.GetDataroomCustomer)
}
