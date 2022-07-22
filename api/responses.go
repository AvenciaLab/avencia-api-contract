package api

type LimitResponse struct {
	Withdrawn float64 `json:"withdrawn"` 
	Max       float64 `json:"max"` 
}

type UserInfoResponse struct {
	Id     string             `json:"id"`
	Wallet map[string]float64 `json:"wallet"`  // key is currency, value is money amount
	Limits map[string]LimitResponse `json:"limits"`  // key is currency, value is limit
}
type VerifiedCodeResponse struct {
	UserInfo UserInfoResponse `json:"user_info"`
}

type CodeResponse struct {
	TransactionCode string `json:"transaction_code"`
	ExpiresAt       int64  `json:"expires_at"`
}
