package controller

import (
	"net/http"

	"go-ec-sample/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{
		service: service.NewProductService(),
	}
}

func (c *ProductController) Index(ctx *gin.Context) {
	products, err := c.service.GetAllProducts()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "error")
		return
	}

	ctx.HTML(http.StatusOK, "product_list.html", gin.H{
		"Products": products,
	})
}
