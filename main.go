package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Name          string    `json:"name"`
	Description   string  `json:"description"`
	Status        string  `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	OnlineDate  time.Time `json:"onlineDate"`
}


func main() {
 router := gin.Default()
 router.Run(":9003")
}