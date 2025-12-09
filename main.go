package main

import (
	"go-ec-sample/controller"
	"go-ec-sample/db"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{})
	r.LoadHTMLGlob("template/*.html")

	productController := controller.NewProductController()
	r.GET("/products", productController.Index)
	r.GET("/products/:id", productController.Show)

	r.Run(":8080") // http://localhost:8080
}
