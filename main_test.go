package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pasDamola/feature-case-study/controllers"
	"github.com/pasDamola/feature-case-study/models"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}
func TestListProductsHandler(t *testing.T) {
	r := SetupRouter()
	r.GET("/products", controllers.ListProductsHandler)
	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var products []models.CatalogProduct
	json.Unmarshal(w.Body.Bytes(), &products)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(products))
}

func TestNewProductHandler(t *testing.T) {
	r := SetupRouter()
	r.POST("/products", controllers.NewProductHandler)
	product := models.CatalogProduct{
	   Name: "Product 4",
	   Description: "Description of Prodcut 4",
	}
	jsonValue, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}