package initializers

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/pasDamola/feature-case-study/models"
	"github.com/robfig/cron/v3"
)


func StartDownloadProductCron() {
	// Set up cron job
	c := cron.New()
	

	// Schedule the cron job to run every day at 3AM
	c.AddFunc("0 3 * * *", generateXMLData)

	c.Start()
}

func generateXMLData() {
	// Query products from the database
    var products []models.CatalogProduct
    if err := DB.Find(&products).Error; err != nil {
        log.Fatalf("Failed to fetch products: %s", err)
    }

    // Generate XML from products
    xmlData, err := xml.MarshalIndent(products, "", "    ")
    if err != nil {
        log.Fatalf("Failed to generate XML: %s", err)
    }

   fmt.Println(string(xmlData))

    log.Println("XML file generated successfully")
}