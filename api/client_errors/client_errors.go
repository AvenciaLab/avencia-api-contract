package client_errors

import "fmt"

type ClientError struct {
	DisplayMessage string `json:"displayMessage"`
	HTTPCode   int    `json:"-"`
}

func (ce ClientError) Error() string {
	return fmt.Sprintf("An error which will be displayed to the client: %v %v", ce.HTTPCode, ce.DisplayMessage)
}

var InvalidCode = ClientError{
	DisplayMessage: "The provided QR Code is invalid.",
	HTTPCode:   400,
}

var InvalidTransactionType = ClientError{
	DisplayMessage: "Invalid transaction type.",
	HTTPCode:   400,
}

var TransactionIdNotProvided = ClientError{
	DisplayMessage: "Transaction ID was not provided.",
	HTTPCode:   400,
}

var IdNotProvided = ClientError{
	DisplayMessage: "ID was not provided in the url query.",
	HTTPCode: 400, 
}

var InvalidATMSecret = ClientError{
	DisplayMessage: "The provided ATM Authentication secret is invalid.",
	HTTPCode:   401,
}

var InsufficientFunds = ClientError{
	DisplayMessage: "Not enough funds to perform the operation.",
	HTTPCode:   400,
}

var WithdrawLimitExceeded = ClientError{
	DisplayMessage: "Withdraw limit exceeded.",
	HTTPCode: 400,
}

var NotFound = ClientError{
	DisplayMessage: "Not Found.", 
	HTTPCode: 404, 
}

var NegativeTransferAmount = ClientError{
	DisplayMessage: "Negative transfer amount.", 
	HTTPCode: 400, 
}

var InvalidJSON = ClientError{
	DisplayMessage: "Invalid request format.", 
	HTTPCode: 400, 
}

var TransferringToYourself = ClientError{
	DisplayMessage: "You cannot transfer to yourself.", 
	HTTPCode: 400, 
}

var TransferringZero = ClientError{
	DisplayMessage: "Transferring zero.", 
	HTTPCode: 400, 
}
