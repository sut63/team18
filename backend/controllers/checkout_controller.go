package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/checkin"
	"github.com/team18/app/ent/checkout"
	"github.com/team18/app/ent/counterstaff"
	"github.com/team18/app/ent/status"
	"github.com/team18/app/ent/statusopinion"
)

// CheckoutController defines the struct for the checkout controller
type CheckoutController struct {
	client *ent.Client
	router gin.IRouter
}

// CheckOut struct
type CheckOut struct {
	CheckoutDate    string
	StatussID       int
	CounterstaffsID int
	CheckinsID      int
	Identitycard    string
	StatusopinionID int
	Comment         string
	Price           float64
}

// CreateCheckout handles POST requests for adding checkout entities
// @Summary Create checkout
// @Description Create checkout
// @ID create-checkout
// @Accept   json
// @Produce  json
// @Param checkout body ent.Checkout true "Checkout entity"
// @Success 200 {object} ent.Checkout
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkouts [post]
func (ctl *CheckoutController) CreateCheckout(c *gin.Context) {
	obj := CheckOut{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "checkout binding failed",
		})
		return
	}

	t1 := time.Now()
	t2 := t1.Format("2006-01-02T15:04:05Z07:00")
	time, err := time.Parse(time.RFC3339, t2)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "time not found",
		})
		return
	}

	cou, err := ctl.client.CounterStaff.
		Query().
		Where(counterstaff.IDEQ(int(obj.CounterstaffsID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "CounterStaff not found",
		})
		return
	}

	so, err := ctl.client.StatusOpinion.
		Query().
		Where(statusopinion.IDEQ(int(obj.StatusopinionID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "StatusopinionID not found",
		})
		return
	}

	st, err := ctl.client.Status.
		Query().
		Where(status.IDEQ(int(1))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Status not found",
		})
		return
	}

	ci, err := ctl.client.CheckIn.
		Query().
		Where(checkin.IDEQ(int(obj.CheckinsID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "CounterStaff not found",
		})
		return
	}

	co, err := ctl.client.Checkout.
		Create().
		SetPrice(obj.Price).
		SetComment(obj.Comment).
		SetIdentityCard(obj.Identitycard).
		SetStatusopinion(so).
		SetCheckoutDate(time).
		SetStatuss(st).
		SetCounterstaffs(cou).
		SetCheckins(ci).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
			"this2":  "this2",
		})
		return
	}
	fmt.Print(err)

	upc, err := ctl.client.CheckIn.
		UpdateOne(ci).
		SetStatusID(2).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "update chackin failed",
		})
		return
	}

	da, err := ctl.client.CheckIn.
		QueryDataroom(ci).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "not found dataroom",
		})
		return
	}
	fmt.Print(err)

	di, err := ctl.client.DataRoom.
		UpdateOne(da).
		SetStatusroomID(1).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "update Dataroom failed",
		})
		return
	}
	fmt.Print(err)

	fmt.Print(upc, di)
	c.JSON(200, gin.H{
		"status": true,
		"data":   co,
	})
}

// ListCheckout handles request to get a list of checkout entities
// @Summary List checkout entities
// @Description list checkout entities
// @ID list-checkout
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Checkout
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkouts [get]
func (ctl *CheckoutController) ListCheckout(c *gin.Context) {
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

	checkouts, err := ctl.client.Checkout.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, checkouts)
}

// GetCheckout handles GET requests to retrieve a checkout entity
// @Summary Get a checkout entity by ID
// @Description get checkout by ID
// @ID get-checkout
// @Produce  json
// @Param id path int true "Checkout ID"
// @Success 200 {array} ent.Checkout
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkouts/{id} [get]
func (ctl *CheckoutController) GetCheckout(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Checkout.
		Query().
		WithCheckins().
		WithStatuss().
		WithStatusopinion().
		WithCounterstaffs().
		Where(checkout.HasStatusopinionWith(statusopinion.IDEQ(int(id)))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// GetCheckout2 handles GET requests to retrieve a GetCheckout2 entity
// @Summary Get a GetCheckout2 entity by ID
// @Description get GetCheckout2 by ID
// @ID get-GetCheckout2
// @Produce  json
// @Param id path int true "Checkout ID"
// @Success 200 {object} ent.ReserveRoom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkouts2/{id} [get]
func (ctl *CheckoutController) GetCheckout2(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.CheckIn.
		Query().
		Where(checkin.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	r, err := ctl.client.CheckIn.
		QueryReserveroom(u).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)
}

// DeleteCheckout handles DELETE requests to delete a checkout entity
// @Summary Delete a checkout entity by ID
// @Description get checkout by ID
// @ID delete-checkout
// @Produce  json
// @Param id path int true "Checkout ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkouts/{id} [delete]
func (ctl *CheckoutController) DeleteCheckout(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Checkout.
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

// NewCheckoutController creates and registers handles for the checkout controller
func NewCheckoutController(router gin.IRouter, client *ent.Client) *CheckoutController {
	uc := &CheckoutController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitCheckoutController registers routes to the main engine
func (ctl *CheckoutController) register() {
	checkouts := ctl.router.Group("/checkouts")
	checkouts2 := ctl.router.Group("/checkouts2")

	checkouts.GET("", ctl.ListCheckout)

	// CRUD
	checkouts.POST("", ctl.CreateCheckout)
	checkouts.GET(":id", ctl.GetCheckout)
	checkouts2.GET(":id", ctl.GetCheckout2)
	checkouts.DELETE(":id", ctl.DeleteCheckout)
}
