package walletapi

import (
	"errors"

	"github.com/shopspring/decimal"
)

// Logic to vlaidate the business logic contratints on the wallet balance and amounts
func ValidateDebitBalance(amount decimal.Decimal, decBalance decimal.Decimal) (bool, error) {

	var allowed bool = false

	if !decBalance.Sub(amount).IsNegative() {

		allowed = true
		return allowed, nil
	} else {
		return allowed, errors.New("error: unable to debit acount, please check your balance")
	}
}

// Non negative vaules only are allowed by this function
func ValidatePositiveAmount(amount decimal.Decimal) (bool, error) {

	var allowed bool = false

	if !amount.IsNegative() {
		allowed = true
		return allowed, nil
	} else {
		return allowed, errors.New("error: Invalid amount, negative values are not allowed")
	}
}
