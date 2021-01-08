package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/counterstaff"
)

//CounterStaffController defines the struct for the course item controller
type CounterStaffController struct {
	client *ent.Client
	router gin.IRouter
}

type CounterStaff struct {
	name     string
	email    string
	password string
}

// CreateCounterStaff handles POST requests for adding CounterStaff entities
// @Summary Create CounterStaff
// @Description Create CounterStaff
// @ID create-CounterStaff
// @Accept   json
// @Produce  json
// @Param CounterStaff body ent.CounterStaff true "CounterStaff entity"
// @Success 200 {object} ent.CounterStaff
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /CounterStaffs [post]
func (ctl *CounterStaffController) CreateCounterStaff(c *gin.Context) {
	obj := CounterStaff{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Counter Staff binding failed",
		})
		return
	}

	csf, err := ctl.client.CounterStaff.
		Create().
		SetName(obj.name).
		SetEmail(obj.email).
		SetPassword(obj.password).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Create Counter Staff error",
		})
		return
	}

	c.JSON(200, csf)
}

// ListCounterStaff handles request to get a list of CounterStaff entities
// @Summary List CounterStaff entities
// @Description list CounterStaff entities
// @ID list-CounterStaff
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.CounterStaff
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /CounterStaffs [get]
func (ctl *CounterStaffController) ListCounterStaff(c *gin.Context) {
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

	customer, err := ctl.client.CounterStaff.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, customer)
}

// GetCounterStaff handles GET requests to retrieve a CounterStaff entity
// @Summary Get a CounterStaff entity by ID
// @Description get CounterStaff by ID
// @ID get-CounterStaff
// @Produce  json
// @Param id path int true "CounterStaff ID"
// @Success 200 {object} ent.CounterStaff
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /CounterStaffs/{id} [get]
func (ctl *CounterStaffController) GetCounterStaff(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.CounterStaff.
		Query().
		Where(counterstaff.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// DeleteCounterStaff handles DELETE requests to delete a CounterStaff entity
// @Summary Delete a CounterStaff entity by ID
// @Description get CounterStaff by ID
// @ID delete-CounterStaff
// @Produce  json
// @Param id path int true "CounterStaff ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /CounterStaffs/{id} [delete]
func (ctl *CounterStaffController) DeleteCounterStaff(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.CounterStaff.
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

// UpdateCounterStaff handles PUT requests to update a CounterStaff entity
// @Summary Update a CounterStaff entity by ID
// @Description update CounterStaff by ID
// @ID update-CounterStaff
// @Accept   json
// @Produce  json
// @Param id path int true "CounterStaff ID"
// @Param CounterStaff body ent.CounterStaff true "CounterStaff entity"
// @Success 200 {object} ent.CounterStaff
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /CounterStaffs/{id} [put]
func (ctl *CounterStaffController) UpdateCounterStaff(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.CounterStaff{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Counter Staff binding failed",
		})
		return
	}
	obj.ID = int(id)

	ci, err := ctl.client.CounterStaff.
		UpdateOne(&obj).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, ci)
}

// NeCounterStaffControllercreates and registers handles for the CounterStaff controller
func NewCounterStaffController(router gin.IRouter, client *ent.Client) *CounterStaffController {
	cs := &CounterStaffController{
		client: client,
		router: router,
	}

	cs.register()

	return cs

}

// InitCounterStaffController registers routes to the main engine
func (ctl *CounterStaffController) register() {
	counterstaffs := ctl.router.Group("/CounterStaffs")

	counterstaffs.GET("", ctl.ListCounterStaff)

	// CRUD
	counterstaffs.POST("", ctl.CreateCounterStaff)
	counterstaffs.PUT(":id", ctl.UpdateCounterStaff)
	counterstaffs.GET(":id", ctl.GetCounterStaff)
	counterstaffs.DELETE(":id", ctl.DeleteCounterStaff)
}
