package api


// ExchangeRatesResponse is a map of currencies and their rates of USD/currency
type ExchangeRatesResponse = map[string]float64 

type OnTransactionCreateResponse struct {
	TransactionId string           `json:"transactionId"`
	Customer      CustomerResponse `json:"customer"`
}

type IdResponse struct {
	Id string `json:"id"`
}

type UserInfoResponse struct {
	Wallets WalletsResponse `json:"wallets"`
	History TransactionHistory `json:"history"`
}

type WalletResponse struct {
	Id       string  `json:"id"`
	OwnerId  string  `json:"owner_id"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type WalletsResponse struct {
	Wallets []WalletResponse `json:"wallets"`
}

type CustomerResponse struct {
	Id        string `json:"id"`
	Email     string `json:"email,omitempty"`
	Mobile    string `json:"mobile,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

type LimitResponse struct {
	Withdrawn float64 `json:"withdrawn"`
	Max       float64 `json:"max"`
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
