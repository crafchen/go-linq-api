package controllers

import (
	"go-linq-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WardController struct {
	service services.WardService
}

func NewWardController(service services.WardService) *WardController {
	return &WardController{service: service}
}

func (c *WardController) RegisterRoutes(r *gin.Engine) {
	g := r.Group("/api/wards")
	{
		g.GET("/", c.GetAll)
		g.GET("/details", c.GetWardDetails)
	}
}

func (c *WardController) GetAll(ctx *gin.Context) {
	wards, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, wards)
}

func (c *WardController) GetWardDetails(ctx *gin.Context) {
	details, err := c.service.GetWardDetails()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, details)
}
