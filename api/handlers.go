package api

import "net/http"

type Handlers struct {
	GenCode, VerifyCode, CheckBanknote, FinalizeTransaction, GetUserInfo http.HandlerFunc
}
