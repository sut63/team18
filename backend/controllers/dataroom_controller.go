package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team18/app/ent"
)

// DataRoomController defines the struct for the course item controller
type DataRoomController struct {
	client *ent.Client
	router gin.IRouter
}

// ListDataRoom handles request to get a list of DataRoom entities
// @Summary List DataRoom entities
// @Description list DataRoom entities
// @ID list-DataRoom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.DataRoom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /DataRooms [get]
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

	reserves, err := ctl.client.DataRoom.
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

	c.JSON(200, reserves)
}

// NewDataRoomControllercreates and registers handles for the DataRoom controller
func NewDataRoomController(router gin.IRouter, client *ent.Client) *DataRoomController {
	dr := &DataRoomController{
		client: client,
		router: router,
	}

	dr.register()

	return dr

}

// InitDataRoomController registers routes to the main engine
func (ctl *DataRoomController) register() {
	DataRooms := ctl.router.Group("/DataRooms")

	DataRooms.GET("", ctl.ListDataRoom)

}
