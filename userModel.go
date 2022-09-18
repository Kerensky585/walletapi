package walletapi

import (
	"gorm.io/gorm"
)

// Struct to represent a very basic user Entity - use to create out schema as well
type user struct {
	gorm.Model
	UID  string  `json:"uid"`
	Name string  `json:"name"`
	Auth float64 `json:"auth"`
}
