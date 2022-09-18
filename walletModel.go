package walletapi

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Struct to represent a very basic wallet Entity - use to create out schema
type wallet struct {
	gorm.Model
	WID     string          `json:"wid"`
	UID     string          `json:"name"`
	Balance decimal.Decimal `json:"balance"`
}
