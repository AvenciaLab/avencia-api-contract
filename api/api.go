package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewAPIRouter(h Handlers, clientAuthMW, atmAuthMW Middleware) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RedirectSlashes)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/transaction", func(r chi.Router) {
			r.Use(atmAuthMW)

			// Request: OnTransactionCreateRequest
			// Response: OnTransactionCreateResponse
			r.Post("/onCreate", h.Transaction.OnCreate)

			// Request: No Request Body
			// Response: No Response Body
			r.Post("/onCancel/{transactionId}", h.Transaction.OnCancel)

			r.Route("/deposit", func(r chi.Router) {
				// Request: BanknoteInsertionRequest
				// Response: No Response Body
				r.Post("/onBanknoteEscrow", h.Transaction.Deposit.OnBanknoteEscrow)

				// Request: BanknoteInsertionRequest
				// Response: No Response Body
				r.Post("/onBanknoteAccepted", h.Transaction.Deposit.OnBanknoteAccepted)

				// Request: CompleteDepositRequest
				// Response: No Response Body
				r.Post("/onComplete", h.Transaction.Deposit.OnComplete)
			})

			r.Route("/withdrawal", func(r chi.Router) {
				// Request: StartWithdrawalRequest
				// Response: No Response Body
				r.Post("/onStart", h.Transaction.Withdrawal.OnStart)

				// Request: BanknoteDispensionRequest
				// Response: No Response Body
				r.Post("/onPreBanknoteDispensed", h.Transaction.Withdrawal.OnPreBanknoteDispensed)

				// Request: BanknoteDispensionRequest
				// Response: No Response Body
				r.Post("/onPostBanknoteDispensed", h.Transaction.Withdrawal.OnPostBanknoteDispensed)

				// Request: CompleteWithdrawalRequest
				// Response: No Response Body
				r.Post("/onComplete", h.Transaction.Withdrawal.OnComplete)
			})

		})

		r.Route("/app", func(r chi.Router) {
			r.Use(clientAuthMW) 
			// Request: GenTransCodeRequest
			// Response: GenTransCodeResponse
			r.Post("/gen-transaction-code", h.App.GenCode)

			// Request: TransferRequest
			// Response: 200 if accepted, client error (or 500) if rejected
			r.Post("/transfer", h.App.Transfer)

			// Request: No Request Body
			// Response: UserInfoResponse
			r.Route("/user", func(r chi.Router) {
				r.Get("/", h.App.GetUserInfo)
				// An endpoint that supports Reading and Updating a DetailedUser entity 
				r.Route("/details", h.App.UserDetails) 
				r.Route("/kyc", func(r chi.Router) {
					r.Route("/passport", h.App.Kyc.Passport)
				})
			})

			// Request: No Request Body
			// Response: TransactionHistoryResponse
			r.Get("/history", h.App.GetUserInfo)
		})

	})

	return r
}

type Middleware = func(http.Handler) http.Handler
