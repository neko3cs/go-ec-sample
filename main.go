package main

import (
	"go-ec-sample/controller"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{})
	r.LoadHTMLGlob("template/*.html")

	productController := controller.NewProductController()
	r.GET("/products", productController.Index)

	r.Run(":8080") // http://localhost:8080
}
