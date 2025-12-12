package controller

import (
	"go-ec-sample/service"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	service *service.LoginService
}

func NewLoginController(s *service.LoginService) *LoginController {
	return &LoginController{service: s}
}

func (lc *LoginController) ShowLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func (c *LoginController) Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	authenticated, user := c.service.Authenticate(email, password)
	if !authenticated {
		ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Error": "Invalid email or password",
		})
		return
	}

	session := sessions.Default(ctx)
	session.Set("user_id", user.Id())
	session.Set("is_admin", user.IsAdmin())
	session.Save()

	ctx.Redirect(http.StatusFound, "/products")
}
