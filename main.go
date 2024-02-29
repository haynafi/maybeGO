package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haynafi/maybeGO/controllers/productcontroller"
	"github.com/haynafi/maybeGO/models"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/haynafi/maybeGO/docs"
)

// @title           CRUD with Gin Golang
// @version         1.0
// @description     wkwkwkwkwk description
// @termsOfService  http://swagger.io/terms/

// @contact.name   saya
// @contact.url    http://github.com/haynafi

// @host      localhost:8080
// @BasePath  /api/

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	//docs with swagger
	r.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}
