package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haynafi/maybeGO/controllers/productcontroller"
	"github.com/haynafi/maybeGO/models"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/haynafi/maybeGO/docs"
)

// @title           CRUD API GIN GOLANG
// @version         1.0
// @description     Sample API with Gin, Golang, and Gorm

// @contact.name   haynafi
// @contact.url    https://github.com/haynafi

// @host      localhost:8080

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	//docs with swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}
