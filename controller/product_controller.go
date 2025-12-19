package controller

import (
	"net/http"
	"strconv"

	"go-ec-sample/consts"
	"go-ec-sample/service"

	"github.com/gin-contrib/sessions"
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

	session := sessions.Default(ctx)
	isAdmin := session.Get(consts.SessionKeyIsAdmin)
	ctx.HTML(http.StatusOK, "product_list.html", gin.H{
		"Products": products,
		"IsAdmin":  isAdmin,
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

	session := sessions.Default(ctx)
	isAdmin := session.Get(consts.SessionKeyIsAdmin)
	ctx.HTML(200, "product_detail.html", gin.H{
		"Product": product,
		"IsAdmin": isAdmin,
	})
}

func (c *ProductController) New(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "product_new.html", gin.H{})
}

func (c *ProductController) Create(ctx *gin.Context) {
	name := ctx.PostForm("name")
	priceStr := ctx.PostForm("price")
	stockStr := ctx.PostForm("stock")

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		ctx.String(400, "Invalid price")
		return
	}
	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		ctx.String(400, "Invalid stock")
		return
	}

	err = c.service.CreateProduct(name, price, stock)
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
	name := ctx.PostForm("name")
	priceStr := ctx.PostForm("price")
	stockStr := ctx.PostForm("stock")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.String(400, "Invalid ID")
		return
	}
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		ctx.String(400, "Invalid price")
		return
	}
	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		ctx.String(400, "Invalid stock")
		return
	}

	err = c.service.UpdateProduct(uint(id), name, price, stock)
	if err != nil {
		ctx.String(500, "Failed to update product")
		return
	}

	ctx.Redirect(302, "/products/"+idStr)
}

func (c *ProductController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.String(400, "Invalid ID")
		return
	}

	err = c.service.DeleteProduct(uint(id))
	if err != nil {
		ctx.String(500, "Failed to delete product")
		return
	}

	ctx.Redirect(302, "/products")
}
