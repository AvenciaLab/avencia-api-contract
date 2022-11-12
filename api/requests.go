package api

type GetExchangeRatesRequest struct {
	Currencies []string `json:"currencies"`
}

type GenTransCodeRequest struct {
	TransactionType string `json:"type"`
	WalletId string `json:"id"`
}

type OnTransactionCreateRequest struct {
	TransactionReference string `json:"transactionReference"`
	TerminalId string `json:"terminalId"`
	TerminalSid string `json:"terminalSid"`
	Type string `json:"type"`
	QRCodeText string `json:"qrCodeText"`
}


type Banknote struct {
	Currency string `json:"currency"`
	Denomination int `json:"denomination"`
}

type BanknoteInsertionRequest struct {
	TransactionId string `json:"transactionId"`
	Banknote Banknote `json:"banknote"`
	Receivables []Money `json:"receivables"`
}

type CompleteDepositRequest struct {
	TransactionId string `json:"transactionId"`
	Receivables []Money `json:"receivables"`
}

type StartWithdrawalRequest struct {
	TransactionId string `json:"transactionId"`
	Currency string `json:"currency"`
	Amount float64 `json:"amount"`
}

type BanknoteDispensionRequest struct {
	TransactionId string `json:"transactionId"`
	Currency string `json:"currency"`
	BanknoteDenomination int `json:"banknoteDenomination"`
	RemainingAmount float64 `json:"remainingAmount"`
	RequestedAmount float64 `json:"requestedAmount"`
}

type CompleteWithdrawalRequest struct {
	TransactionId string `json:"transactionId"`
	Currency string `json:"currency"`
	Amount float64 `json:"amount"`
}



type TransferRequest struct {
	RecipientIdentifier string  `json:"recipient_identifier"` // currently it's email, maybe later it will be a phone number
	Money Money `json:"money"`
}

type CreateWalletRequest struct {
	Currency string `json:"currency"`
}
