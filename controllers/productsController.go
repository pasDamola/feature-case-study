package controllers

import (
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
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func SearchProductsHandler(c *gin.Context) {
	var products []models.CatalogProduct
	initializers.DB.Find(&products)

	tag := c.Query("online_date")
	listOfProducts := make([]models.CatalogProduct, 0)
	for i := 0; i < len(products); i++ {
		found := false
		
		if products[i].OnlineDate != nil && tag == *(products[i].OnlineDate){
			found = true
		}
		if found {
		 listOfProducts = append(listOfProducts, 
			   products[i])
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"products": listOfProducts,
	})
 }