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

func NewProductController(s *service.ProductService) *ProductController {
	return &ProductController{service: s}
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

func (c *ProductController) New(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "product_new.html", gin.H{})
}

func (c *ProductController) Create(ctx *gin.Context) {
	name := ctx.PostForm("name")
	priceStr := ctx.PostForm("price")

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		ctx.String(400, "Invalid price")
		return
	}

	err = c.service.CreateProduct(name, price)
	if err != nil {
		ctx.String(500, "Failed to create product")
		return
	}

	ctx.Redirect(302, "/products")
}

func (c *ProductController) Edit(ctx *gin.Context) {
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

	ctx.HTML(200, "product_edit.html", gin.H{
		"Product": product,
	})
}

func (c *ProductController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.String(400, "Invalid ID")
		return
	}

	name := ctx.PostForm("name")
	priceStr := ctx.PostForm("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		ctx.String(400, "Invalid price")
		return
	}

	err = c.service.UpdateProduct(uint(id), name, price)
	if err != nil {
		ctx.String(500, "Failed to update product")
		return
	}

	ctx.Redirect(302, "/products/"+idStr)
}
