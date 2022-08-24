package client_errors

import "fmt"

type ClientError struct {
	DetailCode string `json:"detail_code"`
	HTTPCode   int    `json:"-"`
}

func (ce ClientError) Error() string {
	return fmt.Sprintf("An error which will be displayed to the client: %v %v", ce.HTTPCode, ce.DetailCode)
}

var InvalidCode = ClientError{
	DetailCode: "invalid-code",
	HTTPCode:   400,
}

var InvalidTransactionType = ClientError{
	DetailCode: "invalid-transaction-type",
	HTTPCode:   400,
}

var TransactionIdNotProvided = ClientError{
	DetailCode: "transaction-id-not-provided",
	HTTPCode:   400,
}

var InvalidATMSecret = ClientError{
	DetailCode: "invalid-atm-secret",
	HTTPCode:   401,
}

var InsufficientFunds = ClientError{
	DetailCode: "insufficient-funds",
	HTTPCode:   400,
}

var WithdrawLimitExceeded = ClientError{
	DetailCode: "withdraw-limit-exceeded",
	HTTPCode: 400,
}

var NotFound = ClientError{
	DetailCode: "not-found", 
	HTTPCode: 404, 
}

var NegativeTransferAmount = ClientError{
	DetailCode: "negative-transfer-amount", 
	HTTPCode: 400, 
}

var InvalidJSON = ClientError{
	DetailCode: "invalid-json", 
	HTTPCode: 400, 
}

var TransferringToYourself = ClientError{
	DetailCode: "transferring-yourself", 
	HTTPCode: 400, 
}

var TransferringZero = ClientError{
	DetailCode: "transferring-zero", 
	HTTPCode: 400, 
}
