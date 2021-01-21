package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
	"github.com/team18/app/ent/statusopinion"
)

// StatusOpinionController defines the struct for the statusopinion controller
type StatusOpinionController struct {
	client *ent.Client
	router gin.IRouter
}

// StatusOpinion struct
type StatusOpinion struct {
	Opinion string
}

// CreateStatusOpinion handles POST requests for adding statusopinion entities
// @Summary Create statusopinion
// @Description Create statusopinion
// @ID create-statusopinion
// @Accept   json
// @Produce  json
// @Param statusopinion body ent.StatusOpinion true "StatusOpinion entity"
// @Success 200 {object} ent.StatusOpinion
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusopinions [post]
func (ctl *StatusOpinionController) CreateStatusOpinion(c *gin.Context) {
	obj := StatusOpinion{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "StatusOpinion binding failed",
		})
		return
	}

	u, err := ctl.client.StatusOpinion.
		Create().
		SetOpinion(obj.Opinion).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"status": true,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   u,
	})
}

// ListStatusOpinion handles request to get a list of statusopinion entities
// @Summary List statusopinion entities
// @Description list statusopinion entities
// @ID list-statusopinion
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.StatusOpinion
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusopinions [get]
func (ctl *StatusOpinionController) ListStatusOpinion(c *gin.Context) {
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

	so, err := ctl.client.StatusOpinion.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, so)
}

// GetStatusOpinion handles GET requests to retrieve a statusopinion entity
// @Summary Get a statusopinion entity by ID
// @Description get statusopinion by ID
// @ID get-statusopinion
// @Produce  json
// @Param id path int true "StatusOpinion ID"
// @Success 200 {object} ent.StatusOpinion
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusopinions/{id} [get]
func (ctl *StatusOpinionController) GetStatusOpinion(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.StatusOpinion.
		Query().
		Where(statusopinion.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// DeleteStatusOpinion handles DELETE requests to delete a statusopinion entity
// @Summary Delete a statusopinion entity by ID
// @Description get statusopinion by ID
// @ID delete-statusopinion
// @Produce  json
// @Param id path int true "StatusOpinion ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusopinions/{id} [delete]
func (ctl *StatusOpinionController) DeleteStatusOpinion(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.StatusOpinion.
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

// NewStatusOpinionController creates and registers handles for the status controller
func NewStatusOpinionController(router gin.IRouter, client *ent.Client) *StatusOpinionController {
	uc := &StatusOpinionController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitStatusOpinionController registers routes to the main engine
func (ctl *StatusOpinionController) register() {
	statusopinions := ctl.router.Group("/statusopinions")

	statusopinions.GET("", ctl.ListStatusOpinion)

	// CRUD
	statusopinions.POST("", ctl.CreateStatusOpinion)
	statusopinions.GET(":id", ctl.GetStatusOpinion)
	statusopinions.DELETE(":id", ctl.DeleteStatusOpinion)
}
