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
	router.LoadHTMLGlob("templates/*.html")
	return router
}

func TestNewProductsHandler(t *testing.T) {
    r := SetupRouter()
    r.POST("/products", controllers.NewProductHandler)
    product := models.CatalogProduct{
        Name: "Demo Name",
        Description: "Demo Description",
    }
    jsonValue, _ := json.Marshal(product)
    req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonValue))

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusCreated, w.Code)
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
    assert.NotEmpty(t, products)
}


func TestAdminPageHandler(t *testing.T) {
    r := SetupRouter()

    r.GET("/admin", controllers.AdminPageHandler)

    req, _ := http.NewRequest("GET", "/admin", nil)


    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    assert.Contains(t, w.Body.String(), "Admin Page")
}

