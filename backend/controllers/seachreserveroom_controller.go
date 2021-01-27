package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/customer"
	"github.com/team18/app/ent/reserveroom"
)

// SeachReserveRoomController defines the struct for the course item controller
type SeachReserveRoomController struct {
	client *ent.Client
	router gin.IRouter
}

// ListReserveCustomee handles request to get a list of ReserveRoom entities by id
// @Summary list a ReserveRoom entity by customerID
// @Description list ReserveRoom entity by customer ID
// @ID list-ReserveCustomer
// @Produce  json
// @Param id path int true "ReserveCustomee ID"
// @Success 200 {array} ent.ReserveRoom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ReserveCustomer/{id} [get]
func (ctl *SeachReserveRoomController) ListReserveCustomer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	i, err := ctl.client.ReserveRoom.
		Query().
		WithRoom().
		WithCustomer().
		WithPromotion().
		WithStatus().
		Where(reserveroom.HasCustomerWith(customer.IDEQ(int(id)))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, i)
}

// NewReserveRoomControllercreates and registers handles for the ReserveRoom controller
func NewSeachReserveRoomController(router gin.IRouter, client *ent.Client) *SeachReserveRoomController {
	sr := &SeachReserveRoomController{
		client: client,
		router: router,
	}

	sr.register()

	return sr

}

// InitSeachReserveRoomController registers routes to the main engine
func (ctl *SeachReserveRoomController) register() {
	ReserveCustomers := ctl.router.Group("/ReserveCustomer")

	// CRUD
	ReserveCustomers.GET(":id", ctl.ListReserveCustomer)
}
