package initializers

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/pasDamola/feature-case-study/models"
	"github.com/robfig/cron/v3"
)


func StartDownloadProductCron() {
	// Set up cron job
	c := cron.New()
	

	// Schedule the cron job to run every day at 3AM
	c.AddFunc(os.Getenv("CRON_EVERY_DAY_AT_3AM"), generateXMLData)

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

	err = os.WriteFile("./tmp/feed/products.xml", xmlData, 0644)
    if err != nil {
        fmt.Println("Error writing XML to file:", err)
        return
    }


    log.Println("XML file generated successfully")
}