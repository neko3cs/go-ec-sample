package main

import (
	"go-ec-sample/controller"
	"go-ec-sample/db"
	"go-ec-sample/service"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{})
	r.LoadHTMLGlob("template/*.html")

	productController := controller.NewProductController(
		service.NewProductService(),
	)
	r.GET("/products", productController.Index)
	r.GET("/products/:id", productController.Show)
	r.GET("/products/new", productController.New)
	r.POST("/products", productController.Create)
	r.GET("/products/:id/edit", productController.Edit)
	r.POST("/products/:id", productController.Update)
	r.POST("/products/:id/delete", productController.Delete)

	r.Run(":8080") // http://localhost:8080
}
