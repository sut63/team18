package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/customer"
)

//CustomerController defines the struct for the course item controller
type CustomerController struct {
	client *ent.Client
	router gin.IRouter
}

// ListCustomer handles request to get a list of Customer entities
// @Summary List Customer entities
// @Description list Customer entities
// @ID list-Customer
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Customer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Customers [get]
func (ctl *CustomerController) ListCustomer(c *gin.Context) {
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

	customer, err := ctl.client.Customer.
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

// GetCustomer handles GET requests to retrieve a Customer entity
// @Summary Get a Customer entity by ID
// @Description get Customer by ID
// @ID get-Customer
// @Produce  json
// @Param id path int true "Customer ID"
// @Success 200 {object} ent.Customer
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Customers/{id} [get]
func (ctl *CustomerController) GetCustomer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Customer.
		Query().
		Where(customer.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// NeCustomerControllercreates and registers handles for the Customer controller
func NewCustomerController(router gin.IRouter, client *ent.Client) *CustomerController {
	c := &CustomerController{
		client: client,
		router: router,
	}

	c.register()

	return c

}

// InitCustomerController registers routes to the main engine
func (ctl *CustomerController) register() {
	Customers := ctl.router.Group("/Customers")

	Customers.GET("", ctl.ListCustomer)

	// CRUD
	Customers.GET(":id", ctl.GetCustomer)

}
