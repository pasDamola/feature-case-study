package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Product struct {
	ProductID           string    `json:"product_id"`
	Name          string    `json:"name"`
	Description   string  `json:"description"`
	Status        string  `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	OnlineDate  string `json:"onlineDate"`
}

var products []Product
func init() {
   products = make([]Product, 0)
}


func NewProductHandler(c *gin.Context) {
	var product Product
   if err := c.ShouldBindJSON(&product); err != nil {
       c.JSON(http.StatusBadRequest, gin.H{
          "error": err.Error()})
       return
   }
   product.ProductID = xid.New().String()
   product.CreatedAt = time.Now()
   product.Status = "new"
   products = append(products, product)
   c.JSON(http.StatusOK, product)
}

func ListProductsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func SearchProductsHandler(c *gin.Context) {
   tag := c.Query("online_date")
   listOfProducts := make([]Product, 0)
   for i := 0; i < len(products); i++ {
       found := false
       if tag == products[i].OnlineDate {
		   found = true
	   }
       if found {
		listOfProducts = append(listOfProducts, 
              products[i])
       }
   }
   c.JSON(http.StatusOK, listOfProducts)
}


func main() {
 router := gin.Default()
 router.POST("/products", NewProductHandler)
 router.GET("/products", ListProductsHandler)
 router.GET("/products/search", SearchProductsHandler)
 router.Run(":9003")
}