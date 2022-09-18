package walletapi

import "github.com/shopspring/decimal"

//Logic to vlaidate the business logic contratints on the wallet balance and amounts
func ValidateDebitBalance(amount decimal.Decimal, decBalance decimal.Decimal) bool {

	var allowed bool = false

	if !decBalance.Sub(amount).IsNegative() {
		allowed = true
	}

	return allowed
}

// Non negative vaules only are allowed by this function
func ValidatePositiveAmount(amount decimal.Decimal) bool {

	return !amount.IsNegative()
}
