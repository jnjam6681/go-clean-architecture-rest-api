package entity

import "gorm.io/gorm"

type Todo struct {
	gorm.Model        // Adds fields 'ID', 'CreatedAt', 'UpdatedAt', 'DeletedAt'
	Name       string `json:"name" gorm:"not null"`
	Completed  bool   `json:"completed"`
}
