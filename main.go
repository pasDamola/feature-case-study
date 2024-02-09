package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pasDamola/feature-case-study/controllers"
	"github.com/pasDamola/feature-case-study/initializers"
)




func init() {
   initializers.LoadEnvVariables()
   initializers.ConnectToDB()
   initializers.ConnectToRedis()
   initializers.ConnectToRabbitMQ()
   initializers.StartNewProductCron()
   initializers.StartDownloadProductCron()

}

func Logger() gin.HandlerFunc {
   return func(c *gin.Context) {
       start := time.Now()

       c.Next()

       // Log request details
       end := time.Now()
       latency := end.Sub(start)
       clientIP := c.ClientIP()
       method := c.Request.Method
       statusCode := c.Writer.Status()
       statusText := http.StatusText(statusCode)
       route := c.FullPath()

      
        log.Printf("[%s] %s %s %s %d %s %v\n", end.Format("2006-01-02 15:04:05"), clientIP, method, route, statusCode, statusText, latency)

      
   }
}



func main() {
   
 router := gin.Default()
 router.LoadHTMLGlob("templates/*.html")


 router.Use(Logger())


 router.GET("/admin", controllers.AdminPageHandler)
 router.POST("/products", controllers.NewProductHandler)
 router.GET("/products", controllers.ListProductsHandler)
 router.GET("/products/search", controllers.SearchProductsHandler)
 router.POST("/products/clear", controllers.ClearCacheHandler)

 router.Run()
}