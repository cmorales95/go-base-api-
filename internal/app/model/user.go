package model

import (
	"time"
)

// User model struct of user
type User struct {
	Base
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Email    string    `json:"email"`
	Birth    time.Time `json:"birth"`
	IsActive bool      `json:"is_active"`
	Google   bool      `json:"google"`
	Facebook bool      `json:"facebook"`
	Twitter  bool      `json:"twitter"`
}

// Users slice of user
type Users []User
