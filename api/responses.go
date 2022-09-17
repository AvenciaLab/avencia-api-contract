package api


type OnTransactionCreateResponse struct {
	TransactionId string `json:"transactionId"`
	Customer CustomerResponse `json:"customer"`
}

type CustomerResponse struct {
	Id string `json:"id"`
	Email string `json:"email,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
}


type LimitResponse struct {
	Withdrawn float64 `json:"withdrawn"`
	Max       float64 `json:"max"`
}

type UserInfoResponse struct {
	Id     string                   `json:"id"`
	Wallet map[string]float64       `json:"wallet"` // key is currency, value is money amount
	Limits map[string]LimitResponse `json:"limits"` // key is currency, value is limit
}

type GenTransCodeResponse struct {
	TransactionCode string `json:"transaction_code"`
	ExpiresAt       int64  `json:"expires_at"` // unix time
}

type TransEntry struct {
	TransactedAt int64             `json:"time"` // unix time
	Source       TransactionSource `json:"source"`
	Money        Money             `json:"money"`
}

type TransactionHistory struct {
	Entries []TransEntry `json:"entries"` // sorted newest-first
}
