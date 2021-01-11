package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/furnituretype"
)

// FurnituretypeController defines the struct for the furnituretype controller
type FurnituretypeController struct {
	client *ent.Client
	router gin.IRouter
}

// GetFurnituretype handles GET requests to retrieve a furnituretype entity
// @Summary Get a furnituretype entity by ID
// @Description get furnituretype by ID
// @ID get-furnituretype
// @Produce  json
// @Param id path int true "Furnituretype ID"
// @Success 200 {object} ent.Furnituretype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /furnituretypes/{id} [get]
func (ctl *FurnituretypeController) GetFurnituretype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	s, err := ctl.client.FurnitureType.
		Query().
		Where(furnituretype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListFurnituretype handles request to get a list of furnituretype entities
// @Summary List furnituretype entities
// @Description list furnituretype entities
// @ID list-furnituretype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Furnituretype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /furnituretypes [get]
func (ctl *FurnituretypeController) ListFurnituretype(c *gin.Context) {
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

	furnituretypes, err := ctl.client.FurnitureType.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, furnituretypes)
}

// NewFurnituretypeController creates and registers handles for the furnituretype controller
func NewFurnituretypeController(router gin.IRouter, client *ent.Client) *FurnituretypeController {
	pc := &FurnituretypeController{
		client: client,
		router: router,
	}
	pc.register()
	return pc
}

// InitFurnituretypeController registers routes to the main engine
func (ctl *FurnituretypeController) register() {
	furnituretypes := ctl.router.Group("/furnituretypes")
	furnituretypes.GET("", ctl.ListFurnituretype)
	// CRUD
	furnituretypes.GET(":id", ctl.GetFurnituretype)
}
