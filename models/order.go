package models

import "time"

type Order struct {
    ID         uint      `gorm:"primaryKey" json:"id"`
    Item       string    `gorm:"size:100;not null" validate:"required" json:"item"`
    Amount     float64   `gorm:"not null" validate:"required,gt=0" json:"amount"`
    Time       time.Time `gorm:"not null" validate:"required" json:"time"`
    CustomerID uint      `gorm:"not null" validate:"required" json:"customerid"`
}
