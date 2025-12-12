package controller

import (
	"bytes"
	"encoding/gob"
	"go-ec-sample/domain"
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

func (cc *CartController) Index(c *gin.Context) {
	session := sessions.Default(c)
	cartInterface := session.Get("cart")

	var cart *domain.Cart
	if cartInterface != nil {
		if bytesData, ok := cartInterface.([]byte); ok {
			buf := bytes.NewBuffer(bytesData)
			var tmpCart domain.Cart
			if err := gob.NewDecoder(buf).Decode(&tmpCart); err == nil {
				cart = &tmpCart
			}
		}
	}
	if cart == nil {
		cart = &domain.Cart{}
	}

	c.HTML(http.StatusOK, "cart.html", gin.H{
		"Cart": cart,
	})
}

func (cc *CartController) Add(c *gin.Context) {
	productIdStr := c.PostForm("product_id")
	quantityStr := c.PostForm("quantity") // 文字列なので strconv.Atoi が必要

	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid product ID")
		return
	}
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid quantity")
		return
	}

	q := query.NewGetProductQuery(uint(productId))
	h := query.NewGetProductQueryHandler()
	product, err := h.Handle(q)
	if err != nil {
		c.String(http.StatusBadRequest, "Product not found")
		return
	}

	session := sessions.Default(c)
	cart := cc.loadCart(session)
	cart.AddItem(*product, quantity)
	cc.saveCart(session, cart)

	c.Redirect(http.StatusFound, "/cart")
}

func (cc *CartController) Remove(ctx *gin.Context) {
	productIdStr := ctx.PostForm("product_id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid product ID")
		return
	}

	session := sessions.Default(ctx)
	cart := cc.loadCart(session)
	cart.RemoveItem(uint(productId))
	cc.saveCart(session, cart)

	ctx.Redirect(http.StatusFound, "/cart")
}

// --------------------------
// ユーティリティー関数
// --------------------------

func (cc *CartController) loadCart(session sessions.Session) *domain.Cart {
	cartInterface := session.Get("cart")
	var cart *domain.Cart

	if cartInterface != nil {
		if bytesData, ok := cartInterface.([]byte); ok {
			buf := bytes.NewBuffer(bytesData)
			var tmpCart domain.Cart
			if err := gob.NewDecoder(buf).Decode(&tmpCart); err == nil {
				cart = &tmpCart
			}
		}
	}

	if cart == nil {
		cart = &domain.Cart{}
	}

	return cart
}

func (cc *CartController) saveCart(session sessions.Session, cart *domain.Cart) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(cart); err != nil {
		panic(err)
	}
	session.Set("cart", buf.Bytes())
	session.Save()
}
