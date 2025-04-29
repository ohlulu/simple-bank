package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/ohlulu/simple-bank/utils"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return utils.IsSupportedCurrency(currency)
	}
	return false
}
