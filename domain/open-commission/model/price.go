package model

type Price struct {
	Amount   float64  `json:"amount" bson:"amount"`
	Currency Currency `json:"currency" bson:"currency"`
}

type Currency string

const (
	CurrencyHKD Currency = "HKD"
	CurrencyTWD Currency = "TWD"
	CurrencyUSD Currency = "USD"
)

func NewCurrency(s string) *Currency {
	var currency Currency
	switch s {
	case "hkd", "HKD":
		currency = CurrencyHKD
		return &currency
	case "twd", "TWD":
		currency = CurrencyTWD
		return &currency
	case "usd", "USD":
		currency = CurrencyUSD
		return &currency
	default:
		return nil
	}
}

func (p Price) GetConvPrice() float64 {
	switch p.Currency {
	case CurrencyUSD:
		return p.Amount
	case CurrencyTWD:
		return p.Amount * 0.035
	case CurrencyHKD:
		return p.Amount * 0.13
	}
	return 0
}

func (c Currency) GetConvPrice(amount float64) float64 {
	switch c {
	case CurrencyUSD:
		return amount
	case CurrencyTWD:
		return amount * 0.035
	case CurrencyHKD:
		return amount * 0.13
	}
	return 0
}