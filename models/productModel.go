package models

import "time"

type CatalogProduct struct {
	ProductID        uint           `gorm:"primaryKey"`
	Name      string
	Description string
	Status string
	OnlineDate *string
	CreatedAt time.Time
}