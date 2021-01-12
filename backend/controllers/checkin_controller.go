package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/checkin"
	"github.com/team18/app/ent/counterstaff"
	"github.com/team18/app/ent/customer"
	"github.com/team18/app/ent/dataroom"
	"github.com/team18/app/ent/reserveroom"
	"github.com/team18/app/ent/statuscheckin"
)

// CheckinController defines the struct for the checkin controller
type CheckinController struct {
	client *ent.Client
	router gin.IRouter
}

// CheckIn struct
type CheckIn struct {
	CheckinDate string
	Statusname  string
	Customer    int
	Counter     int
	Reserveroom int
	Dataroom    int
}

// CreateCheckIn handles POST requests for adding checkin entities
// @Summary Create checkin
// @Description Create checkin
// @ID create-checkin
// @Accept   json
// @Produce  json
// @Param checkin body ent.CheckIn true "CheckIn entity"
// @Success 200 {object} ent.CheckIn
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkins [post]
func (ctl *CheckinController) CreateCheckIn(c *gin.Context) {
	obj := CheckIn{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "checkin binding failed",
		})
		return
	}

	settime := time.Now().Format("2006-01-02T15:04:05Z07:00")
	time, err := time.Parse(time.RFC3339, settime)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "time not found",
		})
		return
	}

	cus, err := ctl.client.Customer.
		Query().
		Where(customer.IDEQ(int(obj.Customer))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "customer not found",
		})
		return
	}

	cou, err := ctl.client.CounterStaff.
		Query().
		Where(counterstaff.IDEQ(int(obj.Counter))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "CounterStaff not found",
		})
		return
	}

	r, err := ctl.client.ReserveRoom.
		Query().
		Where(reserveroom.IDEQ(int(obj.Reserveroom))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "CounterStaff not found",
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

	ch, err := ctl.client.CheckIn.
		Create().
		SetCheckinDate(time).
		SetCustomer(cus).
		SetCounter(cou).
		SetReserveroom(r).
		SetDataroom(dr).
		SetStatusID(1).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Create CheckIn error",
		})
		return
	}

	ci, err := ctl.client.ReserveRoom.
		UpdateOne(r).
		SetStatusID(2).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Update status re CheckIn error",
		})
		return
	}
	fmt.Print(ci)

	c.JSON(200, ch)
}

// ListCheckIn handles request to get a list of checkin entities
// @Summary List checkin entities
// @Description list checkin entities
// @ID list-checkin
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.CheckIn
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkins [get]
func (ctl *CheckinController) ListCheckIn(c *gin.Context) {
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

	checkins, err := ctl.client.CheckIn.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, checkins)
}

// GetCheckIn handles GET requests to retrieve a checkin entity
// @Summary Get a checkin entity by ID
// @Description get checkin by ID
// @ID get-checkin
// @Produce  json
// @Param id path int true "Checkin ID"
// @Success 200 {object} ent.CheckIn
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkins/{id} [get]
func (ctl *CheckinController) GetCheckIn(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ch, err := ctl.client.CheckIn.
		Query().
		Where(checkin.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ch)
}

// GetCheckInStatus handles request to get a list of  GetCheckInStatus entities
// @Summary List  GetCheckInStatus entities
// @Description list  GetCheckInStatus entities
// @ID list-GetCheckInStatus
// @Produce json
// @Success 200 {array} ent.CheckIn
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkinstatuss [get]
func (ctl *CheckinController) GetCheckInStatus(c *gin.Context) {

	ch, err := ctl.client.CheckIn.
		Query().
		Where(checkin.HasStatusWith(statuscheckin.StatusNameEQ("พักอยู่"))).
		All(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ch)
}

// DeleteCheckIn handles DELETE requests to delete a checkin entity
// @Summary Delete a checkin entity by ID
// @Description get checkin by ID
// @ID delete-checkin
// @Produce  json
// @Param id path int true "Checkin ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkins/{id} [delete]
func (ctl *CheckinController) DeleteCheckIn(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.CheckIn.
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

// UpdateCheckIn handles PUT requests to update a Checkin entity
// @Summary Update a Checkin entity by ID
// @Description update Checkin by ID
// @ID update-Checkin
// @Accept   json
// @Produce  json
// @Param id path int true "Checkin ID"
// @Param Checkin body ent.CheckIn true "CheckIn entity"
// @Success 200 {object} ent.CheckIn
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkins/{id} [put]
func (ctl *CheckinController) UpdateCheckIn(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.CheckIn{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Checkin binding failed",
		})
		return
	}
	obj.ID = int(id)

	ci, err := ctl.client.CheckIn.
		UpdateOne(&obj).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, ci)
}

// NewCheckinController creates and registers handles for the checkin controller
func NewCheckinController(router gin.IRouter, client *ent.Client) *CheckinController {
	cc := &CheckinController{
		client: client,
		router: router,
	}
	cc.register()
	return cc
}

// InitCheckinController registers routes to the main engine
func (ctl *CheckinController) register() {
	checkins := ctl.router.Group("/checkins")
	checkinstatuss := ctl.router.Group("/checkinstatuss")

	checkins.GET("", ctl.ListCheckIn)

	// CRUD
	checkins.POST("", ctl.CreateCheckIn)
	checkins.PUT(":id", ctl.UpdateCheckIn)
	checkins.GET(":id", ctl.GetCheckIn)
	checkinstatuss.GET("", ctl.GetCheckInStatus)
	checkins.DELETE(":id", ctl.DeleteCheckIn)
}
