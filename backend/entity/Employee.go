package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid:"required~name not blank"`
	Email      string
	EmployeeID string
}
