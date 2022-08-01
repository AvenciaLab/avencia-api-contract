package api

// Requested from the ATM or WhiteEdge
type CodeRequest struct {
	TransactionCode string `json:"transaction_code"`
}
type BanknoteCheckRequest struct {
	TransactionCode string  `json:"transaction_code"`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
}
type FinalizeTransactionRequest struct {
	UserId    string            `json:"user_id"`
	ATMSecret string            `json:"atm_secret"`
	Source    TransactionSource `json:"source"`
	// TODO: change to use Money struct
	Currency  string            `json:"currency"`
	Amount    float64           `json:"amount"` // negative value means withdrawal
}

// Requested from the Avencia App
type TransferRequest struct {
	RecipientIdentifier string  `json:"recipient_identifier"` // currently it's email, maybe later it will be a phone number
	Currency            string  `json:"currency"`
	Amount              float64 `json:"amount"` // can only be positive for obvious reasons
}
