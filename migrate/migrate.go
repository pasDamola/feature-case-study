package main

import (
	"github.com/pasDamola/feature-case-study/initializers"
	"github.com/pasDamola/feature-case-study/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}


func main() {
	initializers.DB.AutoMigrate(&models.CatalogProduct{})
}