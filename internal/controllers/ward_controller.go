package controllers

import (
	"go-linq-api/internal/helpers"
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

// ---------------- GET ALL ----------------
func (c *WardController) GetAll(ctx *gin.Context) {
	result := c.service.GetAll()
	ctx.JSON(http.StatusOK, result)
}

// ---------------- GET WARD DETAILS WITH PAGINATION ----------------
func (c *WardController) GetWardDetails(ctx *gin.Context) {
	var pagination helpers.PaginationParam

	// Bind từ JSON body vào struct
	if err := ctx.ShouldBindJSON(&pagination); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewOperationResultError("Invalid body JSON"))
		return
	}

	pagination.Normalize()
	result := c.service.GetWardDetails(pagination)
	ctx.JSON(http.StatusOK, result)
}
