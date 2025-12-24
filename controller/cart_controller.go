package controller

import (
	"go-ec-sample/consts"
	"go-ec-sample/service"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CartController struct {
	cartService    *service.CartService
	productService *service.ProductService
}

func NewCartController(cartService *service.CartService, productService *service.ProductService) *CartController {
	return &CartController{cartService: cartService, productService: productService}
}

func (c *CartController) Index(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userId := session.Get(consts.SessionKeyUserID).(uint)
	cart, err := c.cartService.GetCart(userId)
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

	session := sessions.Default(ctx)
	userId := session.Get(consts.SessionKeyUserID).(uint)
	err = c.cartService.AddToCart(userId, uint(productId), quantity)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed to add item to cart")
		return
	}
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
	userId := session.Get(consts.SessionKeyUserID).(uint)
	err = c.cartService.RemoveFromCart(userId, uint(productId))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed to remove item from cart")
		return
	}

	ctx.Redirect(http.StatusFound, "/cart")
}

func (c *CartController) Checkout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userId := session.Get(consts.SessionKeyUserID).(uint)
	err := c.cartService.Checkout(userId)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	session.Delete(consts.SessionKeyCart)
	session.Save()
	ctx.Redirect(http.StatusFound, "/products")
}
