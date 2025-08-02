package controllers

import (
	"go-linq-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProvinceController struct {
	service services.ProvinceService
}

func NewProvinceController(service services.ProvinceService) *ProvinceController {
	return &ProvinceController{service: service}
}

func (c *ProvinceController) RegisterRoutes(r *gin.Engine) {
	g := r.Group("/api/provinces")
	{
		g.GET("/", c.GetAll)
		g.GET("/:code", c.GetByCode)
		g.GET("/stats", c.GetWithStatistics)
	}
}

func (c *ProvinceController) GetAll(ctx *gin.Context) {
	provinces, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, provinces)
}

func (c *ProvinceController) GetByCode(ctx *gin.Context) {
	code := ctx.Param("code")
	province, err := c.service.GetByCode(code)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Province not found"})
		return
	}
	ctx.JSON(http.StatusOK, province)
}

func (c *ProvinceController) GetWithStatistics(ctx *gin.Context) {
	stats, err := c.service.GetWithStatistics()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, stats)
}
