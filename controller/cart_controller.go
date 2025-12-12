package controller

import (
	"go-ec-sample/query"
	"go-ec-sample/service"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CartController struct {
	service *service.CartService
}

func NewCartController(s *service.CartService) *CartController {
	return &CartController{service: s}
}

func (c *CartController) Index(ctx *gin.Context) {
	session := sessions.Default(ctx)
	cart, err := c.service.LoadCart(session)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed to load cart")
		return
	}
	ctx.HTML(http.StatusOK, "cart.html", gin.H{
		"Cart": cart,
	})
}

func (c *CartController) Add(ctx *gin.Context) {
	productIdStr := ctx.PostForm("product_id")
	quantityStr := ctx.PostForm("quantity")

	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid product id")
		return
	}
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid quantity")
		return
	}

	q := query.NewGetProductQuery(uint(productId))
	h := query.NewGetProductQueryHandler()
	product, err := h.Handle(q)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Product not found")
		return
	}
	if product.Stock() < quantity {
		ctx.String(http.StatusBadRequest, "Not enough stock")
		return
	}

	session := sessions.Default(ctx)
	cart, err := c.service.LoadCart(session)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed to load cart")
		return
	}
	cart.AddItem(*product, quantity)
	c.service.SaveCart(session, cart)

	ctx.Redirect(http.StatusFound, "/cart")
}

func (c *CartController) Remove(ctx *gin.Context) {
	productIdStr := ctx.PostForm("product_id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid product id")
		return
	}

	session := sessions.Default(ctx)
	cart, err := c.service.LoadCart(session)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed to load cart")
		return
	}
	cart.RemoveItem(uint(productId))
	c.service.SaveCart(session, cart)

	ctx.Redirect(http.StatusFound, "/cart")
}

func (c *CartController) Checkout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	cart, err := c.service.LoadCart(session)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed to load cart")
		return
	}

	err = c.service.Checkout(cart)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	session.Delete("cart")
	session.Save()
	ctx.Redirect(http.StatusFound, "/products")
}
