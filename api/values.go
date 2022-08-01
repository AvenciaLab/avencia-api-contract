package api

type Money struct {
	Currency string  `json:"currency"` // currently accepts any non-empty string
	Amount   float64 `json:"amount"`   // can be both positive and negative (means withdrawal)
}

const (
	CreditCardType = "credit-card"
	CashType       = "cash"
	CryptoType     = "crypto"
	TransferType   = "transfer"
)

// TransactionSource Type must be one of the values defined above
// TransactionSource Detail is:
// - credit card number for the credit-card Type
// - nothing for the cash Type
// - crypto wallet for the crypto Type
// - ID of the sender or the reciever (defined contextually) for the transfer Type
type TransactionSource struct {
	Type   string `json:"type"`
	Detail string `json:"detail"`
}
