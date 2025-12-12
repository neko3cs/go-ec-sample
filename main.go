package main

import (
	"go-ec-sample/controller"
	"go-ec-sample/db"
	"go-ec-sample/middleware"
	"go-ec-sample/service"
	"text/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{})
	r.LoadHTMLGlob("template/*.html")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	loginController := controller.NewLoginController(
		service.NewLoginService(),
	)
	productController := controller.NewProductController(
		service.NewProductService(),
	)

	r.GET("/login", loginController.ShowLogin)
	r.POST("/login", loginController.Login)
	r.GET("/logout", loginController.Logout)
	authed := r.Group("/products")
	authed.Use(middleware.AuthRequired())
	{
		authed.GET("/", productController.Index)
		authed.GET("/:id", productController.Show)
	}
	admin := r.Group("/admin")
	admin.Use(middleware.AuthRequired(), middleware.AdminRequired())
	{
		admin.GET("/products/new", productController.New)
		admin.POST("/products", productController.Create)
		admin.GET("/products/:id/edit", productController.Edit)
		admin.POST("/products/:id", productController.Update)
		admin.POST("/products/:id/delete", productController.Delete)
	}

	r.Run(":8080") // http://localhost:8080/login
}
