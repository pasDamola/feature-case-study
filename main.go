package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pasDamola/feature-case-study/controllers"
	"github.com/pasDamola/feature-case-study/initializers"
)


func init() {
   initializers.LoadEnvVariables()
   initializers.ConnectToDB()
   initializers.ConnectToRedis()
}




func main() {
   
 router := gin.Default()
 router.POST("/products", controllers.NewProductHandler)
 router.GET("/products", controllers.ListProductsHandler)
 router.GET("/products/search", controllers.SearchProductsHandler)
 router.GET("/products/clear", controllers.ClearCacheHandler)
 router.Run()
}