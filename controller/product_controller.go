package controller

import (
	"net/http"
	"strconv"

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

func (c *ProductController) Show(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.String(400, "Invalid ID")
		return
	}

	product, err := c.service.GetProduct(uint(id))
	if err != nil {
		ctx.String(404, "Product not found")
		return
	}

	ctx.HTML(200, "product_detail.html", gin.H{
		"Product": product,
	})
}
