package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
)

//PromotionController defines the struct for the course item controller
type PromotionController struct {
	client *ent.Client
	router gin.IRouter
}

// ListPromotion handles request to get a list of Promotion entities
// @Summary List Promotion entities
// @Description list Promotion entities
// @ID list-Promotion
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Promotion
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Promotions [get]
func (ctl *PromotionController) ListPromotion(c *gin.Context) {
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

	promo, err := ctl.client.Promotion.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, promo)
}

// NewPromotionControllercreates and registers handles for the Promotion controller
func NewPromotionController(router gin.IRouter, client *ent.Client) *PromotionController {
	c := &PromotionController{
		client: client,
		router: router,
	}

	c.register()

	return c

}

// InitPromotionController registers routes to the main engine
func (ctl *PromotionController) register() {
	Promotions := ctl.router.Group("/Promotions")

	Promotions.GET("", ctl.ListPromotion)

}
