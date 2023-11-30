package main

import (
	"github.com/bagaskaramadhan/go-project-toko/controllers/productController"
	"github.com/bagaskaramadhan/go-project-toko/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// create routes
	routes := gin.Default();

	models.ConnectDB();

	routes.GET("/api/v1/product", productController.GetAll);
	routes.GET("/api/v1/product/:id", productController.GetById);
	routes.POST("/api/v1/product", productController.Create);
	routes.PUT("/api/v1/product/:id", productController.Update);
	routes.DELETE("/api/v1/product/:id", productController.Delete);

	// Run the routes
	routes.Run();
}
