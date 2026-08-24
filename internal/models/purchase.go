package models

import "time"

type Purchase struct {
	ID          uint      `gorm:"primaryKey"`
	SupplierID  uint
	Material    string
	Quantity    int
	UnitPrice   float64
	TotalPrice  float64
	PurchaseDate time.Time
	ExpenseType string
	UserID      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
