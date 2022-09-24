package api

import "net/http"

type Endpoint = func(r chi.Router)

type Handlers struct {
	Transaction TransactionHandlers 
	App AppHandlers
}

type TransactionHandlers struct {
	OnCreate, OnCancel http.HandlerFunc

	Deposit    TransactionDepositHandlers
	Withdrawal TransactionWithdrawalHandlers
}

type TransactionDepositHandlers struct {
	OnBanknoteEscrow, OnBanknoteAccepted, OnComplete http.HandlerFunc
}

type TransactionWithdrawalHandlers struct {
	OnStart, OnPreBanknoteDispensed, OnPostBanknoteDispensed, OnComplete http.HandlerFunc 
}

type AppHandlers struct {
	GenCode, GetUserInfo, Transfer, GetHistory http.HandlerFunc
	UserDetails Endpoint 
}
