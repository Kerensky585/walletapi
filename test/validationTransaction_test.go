package test

import (
	"testing"

	"github.com/Kerensky585/walletapi"
	"github.com/shopspring/decimal"
)

func TestInvalidValidateDebitBalance(t *testing.T) {
	result, err := walletapi.ValidateDebitBalance(decimal.New(100.00, 2), decimal.New(20.00, 2))

	if result == true || err == nil {
		t.Log("The result should be false if invalid transaction and we should have an error", err)
		t.Fail()
	}
}

func TestValidValidateDebitBalance(t *testing.T) {
	result, err := walletapi.ValidateDebitBalance(decimal.New(20.00, 2), decimal.New(120.00, 2))

	if result == false || err != nil {
		t.Log("The result should be True for valid transaction and we should have NO error", err)
		t.Fail()
	}
}

func TestInvalidValidatePositiveAmount(t *testing.T) {
	result, err := walletapi.ValidatePositiveAmount(decimal.New(-100.00, 2))

	if result == true || err == nil {
		t.Log("For a Negative amount the result should be FALSE with and we should have an ERROR", err)
		t.Fail()
	}
}

func TestValidValidatePositiveAmount(t *testing.T) {
	result, err := walletapi.ValidatePositiveAmount(decimal.New(100.00, 2))

	if result == false || err != nil {
		t.Log("For a valid positive amount the result should be TRUE with no error", err)
		t.Fail()
	}
}
