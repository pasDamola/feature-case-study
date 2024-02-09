package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pasDamola/feature-case-study/initializers"
	"github.com/pasDamola/feature-case-study/models"
)



func NewProductHandler(c *gin.Context) {
	var body struct {
		Name string
		Description string
		Status string
	}

	c.Bind(&body)

	product := models.CatalogProduct{Name: body.Name, Description: body.Description, Status: "new", OnlineDate: nil}

	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})

}
func ListProductsHandler(c *gin.Context) {
	var products []models.CatalogProduct
	initializers.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func SearchProductsHandler(c *gin.Context) {
	var products []models.CatalogProduct
	initializers.DB.Find(&products)

	online_date := c.Query("online_date")

	// Check if data exists in Redis cache
    cachedData, err := initializers.RedisClient.Get(c, fmt.Sprintf("product-%s", online_date)).Result()

	
	// Convert data from string back to JSON format
	if len(cachedData) > 0 {
		json_err := json.Unmarshal([]byte(cachedData), &products)
		if json_err != nil {
			fmt.Println(json_err)
			return
		}
	


    if err == nil {
        // Data found in cache, return it
        c.JSON(200, gin.H{
            "products": products,
        })
        return
    }

}
	
	listOfProducts := make([]models.CatalogProduct, 0)
	for i := 0; i < len(products); i++ {
		found := false

		
		if products[i].OnlineDate != nil && online_date == *(products[i].OnlineDate){
			found = true
		}
		if found {
		 listOfProducts = append(listOfProducts, 
			   products[i])
		}
	}


	data, _ := json.Marshal(listOfProducts)

	// Save to redis cache
	err = initializers.RedisClient.Set(c, fmt.Sprintf("product-%s", online_date), string(data), 0).Err()
	fmt.Println(err)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to save data to cache"})
        return
    }
	c.JSON(http.StatusOK, gin.H{
		"products": listOfProducts,
	})
 }

 func AdminPageHandler(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"successMessage": "",
		},
	  )
 }

 func ClearCacheHandler(c *gin.Context) {
	err := initializers.RedisClient.FlushAll(c).Err()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to clear cache"})
	}

	successMessage := "Cache cleared successfully"

	c.HTML(http.StatusOK, "index.html", gin.H{
        "successMessage": successMessage,
    })
 }
