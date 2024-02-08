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
   initializers.ConnectToRabbitMQ()
   initializers.StartProductCron()

}

// type Request struct {
//    URL string `json:"url"`
//  }


// func ParserHandler(c *gin.Context) {
   
//    var request Request
//    if err := c.ShouldBindJSON(&request); err != nil {
//        c.JSON(http.StatusBadRequest, gin.H{
//           "error": err.Error()})
//        return
//    }
//    data, _ := json.Marshal(request)
//    err := initializers.ChannelAmqp.Publish(
//        "",
//        os.Getenv("RABBITMQ_QUEUE"),
//        false,
//        false,
//        amqp.Publishing{
//            ContentType: "application/json",
//            Body:        []byte(data),
//        })
//    if err != nil {
//        fmt.Println(err)
//        c.JSON(http.StatusInternalServerError, 
//           gin.H{"error": "Error while publishing to RabbitMQ"})
//        return
//    }
//    c.JSON(http.StatusOK, map[string]string{
//       "message": "success"})
// }


func main() {
   
 router := gin.Default()
 router.LoadHTMLGlob("templates/*.html")



 router.GET("/admin", controllers.AdminPageHandler)
 router.POST("/products", controllers.NewProductHandler)
 router.GET("/products", controllers.ListProductsHandler)
 router.GET("/products/search", controllers.SearchProductsHandler)
 router.POST("/products/clear", controllers.ClearCacheHandler)

 router.Run()
}