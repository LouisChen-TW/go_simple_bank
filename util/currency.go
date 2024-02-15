package util

var currencyMap = map[string]string{
	"USD": "USD",
	"EUR": "EUR",
	"CAD": "CAD",
}

func IsValidCurrency(currency string) bool {
	return currencyMap[currency] != ""
}
