package initializers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pasDamola/feature-case-study/models"
	"github.com/robfig/cron/v3"
	"github.com/streadway/amqp"
)



func StartNewProductCron() {
	 // Set up cron job
	 c := cron.New()
	 

	 // Schedule the cron job to run every 10 minutes
	 c.AddFunc(os.Getenv("CRON_EVERY_10_MINUTES"), publishNewProducts)
 
	 c.Start()
}

func queryNewProducts() []models.CatalogProduct {
    var products []models.CatalogProduct
    result := DB.Where("status = ?", "new").Find(&products)

	if result.Error != nil {
        log.Println("failed to retrieve products with new status")
    }

	return products
}

func publishNewProducts() {
	for _, product := range queryNewProducts() {
        publishToRabbitMQ(product)
        markAsInProgress(product)
    }

	consumeNewProducts()
}

func markAsInProgress(product models.CatalogProduct) {
	result := DB.Model(&models.CatalogProduct{}).Where("product_id = ? AND status = ?", product.ProductID, "new").Update("status", "in-progress")
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

    fmt.Println(os.Getenv("RABBITMQ_QUEUE"))
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


func consumeNewProducts() {
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

    time.Sleep(1 * time.Second)

    // Consume messages from the queue
    msgs, err := ch.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Failed to consume messages: %s", err)
    }

    // Process messages
    for msg := range msgs {
        // Simulate processing time

        // Decode message into Product
        var product models.CatalogProduct
        err := json.Unmarshal(msg.Body, &product)
        if err != nil {
            log.Printf("Failed to decode message: %s", err)
            continue
        }

		currentTime := time.Now()
        currentTimeString := currentTime.Format("2006-01-02 15:04:05")

        // Update online date column
        product.OnlineDate = &currentTimeString

        // Mark status as processed
        product.Status = "processed"

        // Update record in the database
        err = DB.Save(&product).Error
        if err != nil {
            log.Printf("Failed to update record in database: %s", err)
            continue
        }

        log.Printf("Processed product with ID %d", product.ProductID)
    }
}