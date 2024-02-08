package initializers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/pasDamola/feature-case-study/models"
	"github.com/robfig/cron/v3"
	"github.com/streadway/amqp"
)

var c *cron.Cron

func StartProductCron() {
	 // Set up cron job
	 c = cron.New()

	 // Schedule the cron job to run every day at 2:00 AM
	 c.AddFunc("* * * * *", processNewRecords)
 
	 c.Start()
}

func queryNewProducts() []models.CatalogProduct {
    var products []models.CatalogProduct
    result := DB.Where("status = ?", "new").Find(&products)

	if result.Error != nil {
        panic("failed to retrieve products with new status")
    }

	return products
}

func processNewRecords() {
	for _, product := range queryNewProducts() {
		fmt.Println("product", product)
        publishToRabbitMQ(product)
        markAsInProgress(product)
    }
}

func markAsInProgress(product models.CatalogProduct) {
	result := DB.Model(&models.CatalogProduct{}).Where("product_id = ? AND status = ?", product.ProductID, "new").Update("status", "in-progress")
	fmt.Println(result.Error)
	if result.Error != nil {
        log.Fatal("failed to update status")
    }
}

func publishToRabbitMQ(product models.CatalogProduct) {
    // Connect to RabbitMQ
    conn, err := amqp.Dial(os.Getenv("RABBITMQ_URI"))
    if err != nil {
        fmt.Println("Failed to connect to RabbitMQ:", err)
        return
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        fmt.Println("Failed to open channel:", err)
        return
    }
    defer ch.Close()

    
    q, err := ch.QueueDeclare(
        os.Getenv("RABBITMQ_QUEUE"),
        true,                        
        false,                        
        false,                        
        false,                        
        nil,                          
    )
    if err != nil {
        
        fmt.Println("Failed to declare queue:", err)
        return
    }

    
    data, err := json.Marshal(product)
    if err != nil {
        // Handle error
        fmt.Println("Failed to serialize record to JSON:", err)
        return
    }

    
    err = ch.Publish(
        "", 
        q.Name,
        false,
        false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        data,
        },
    )
    if err != nil {
        
        fmt.Println("Failed to publish message to RabbitMQ:", err)
        return
    }

    fmt.Println("Record pushed to RabbitMQ")
}
