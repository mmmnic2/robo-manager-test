package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Robot struct {
	ID                uint64          `gorm:"primary_key;auto_increment" json:"id"`
	Name              string          `gorm:"size:255;not null" json:"name"`
	Description       string          `gorm:"size:255;" json:"description"`
	Price             decimal.Decimal `gorm:"type:decimal(10,2);" json:"price"`
	Model             string          `gorm:"size:255;" json:"model"`
	Version           string          `gorm:"size:255;" json:"version"`
	ManufacturingDate time.Time       `gorm:"type:date;" json:"manufacturing_date"`
	Color             string          `gorm:"size:255;" json:"color"`
	Camera            string          `gorm:"size:255;" json:"camera"`
	BatteryLife       string          `gorm:"size:255;" json:"battery_life"`
	Material          string          `gorm:"size:255;" json:"material"`
	Weight            float32         `gorm:"type:float4;" json:"weight"`
	Speed             float32         `gorm:"type:float4;" json:"speed"`
	Sensors           string          `gorm:"size:255;" json:"sensors"`
	Connectivity      string          `gorm:"size:255;" json:"connectivity"`
}
