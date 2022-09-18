package walletapi

import "github.com/shopspring/decimal"

// Struct to represent a very basic transaction Entity - use to create out schema
type updateBal struct {
	WID    string          `json:"wid"`
	Amount decimal.Decimal `json:"amount"`
}
