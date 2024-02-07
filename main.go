package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pasDamola/feature-case-study/controllers"
	"github.com/pasDamola/feature-case-study/initializers"
)


func init() {
   initializers.LoadEnvVariables()
   initializers.ConnectToDB()
   // products = make([]Product, 0)
}


// func NewProductHandler(c *gin.Context) {
// 	var product Product
//    if err := c.ShouldBindJSON(&product); err != nil {
//        c.JSON(http.StatusBadRequest, gin.H{
//           "error": err.Error()})
//        return
//    }
//    product.ProductID = xid.New().String()
//    product.CreatedAt = time.Now()
//    product.Status = "new"
//    products = append(products, product)
//    c.JSON(http.StatusOK, product)
// }

// func ListProductsHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, products)
// }

// func SearchProductsHandler(c *gin.Context) {
//    tag := c.Query("online_date")
//    listOfProducts := make([]Product, 0)
//    for i := 0; i < len(products); i++ {
//        found := false
//        if tag == products[i].OnlineDate {
// 		   found = true
// 	   }
//        if found {
// 		listOfProducts = append(listOfProducts, 
//               products[i])
//        }
//    }
//    c.JSON(http.StatusOK, listOfProducts)
// }


func main() {
 router := gin.Default()
 router.POST("/products", controllers.NewProductHandler)
 router.GET("/products", controllers.ListProductsHandler)
 router.GET("/products/search", controllers.SearchProductsHandler)
 router.Run()
}