package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/furniture"
)

// FurnitureController defines the struct for the furniture controller
type FurnitureController struct {
	client *ent.Client
	router gin.IRouter
}

// GetFurniture handles GET requests to retrieve a furniture entity
// @Summary Get a furniture entity by ID
// @Description get furniture by ID
// @ID get-furniture
// @Produce  json
// @Param id path int true "Furniture ID"
// @Success 200 {object} ent.Furniture
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /furnitures/{id} [get]
func (ctl *FurnitureController) GetFurniture(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	s, err := ctl.client.Furniture.
		Query().
		Where(furniture.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListFurniture handles request to get a list of furniture entities
// @Summary List furniture entities
// @Description list furniture entities
// @ID list-furniture
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Furniture
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /furnitures [get]
func (ctl *FurnitureController) ListFurniture(c *gin.Context) {
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

	furnitures, err := ctl.client.Furniture.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, furnitures)
}

// NewFurnitureController creates and registers handles for the furniture controller
func NewFurnitureController(router gin.IRouter, client *ent.Client) *FurnitureController {
	pc := &FurnitureController{
		client: client,
		router: router,
	}
	pc.register()
	return pc
}

// InitFurnitureController registers routes to the main engine
func (ctl *FurnitureController) register() {
	furnitures := ctl.router.Group("/furnitures")
	furnitures.GET("", ctl.ListFurniture)
	// CRUD
	furnitures.GET(":id", ctl.GetFurniture)
}
