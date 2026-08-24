package models

import "time"

type Supplier struct {
	ID         uint       `gorm:"primaryKey"`
	Name       string
	Contact    string
	Purchases  []Purchase `gorm:"foreignKey:SupplierID"`
	UserID     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
