package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

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
	GenCode, Transfer, GetHistory, GetUserInfo, GetExchangeRates http.HandlerFunc
	Kyc KycHandlers
	UserDetails Endpoint 
	Address Endpoint
	Wallets WalletHandlers
}

type KycHandlers struct {
	Passport, NationalId, DrivingLicense Endpoint
}

type WalletHandlers struct {
	GetAll, Create http.HandlerFunc 
}
