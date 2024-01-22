package v1

import (
	"kreditplus/src/middleware"
	"kreditplus/src/v1/handler"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router, deps *Dependency) {
	// Check Health
	r.Get("/health", handler.Health())

	// Customer
	r.Post("/customer/login", handler.LoginCustomer(deps.Services.customer))
	r.Route("/customer", func(r chi.Router) {
		r.Use(middleware.ValidateAuthorization())
		r.Use(middleware.CheckAuthorization())

		r.Get("/limit", handler.LimitCustomer(deps.Services.customer))
	})

	// Partner Endpoint
	r.Route("/partner", func(r chi.Router) {
		r.Post("/transaction/init", handler.CreateTransaction(deps.Services.transaction))
	})

	// Transaction
	r.Route("/transaction", func(r chi.Router) {
		r.Use(middleware.ValidateAuthorization())
		r.Use(middleware.CheckAuthorization())

		r.Get("/{txn_id}", handler.GetTransaction(deps.Services.transaction))
		r.Post("/settlement", handler.SettlementTransaction(deps.Services.transaction))
	})

	// Installment
	// Public API and User auth
	r.Route("/installment", func(r chi.Router) {
		r.Use(middleware.ValidateAuthorization())
		r.Post("/calculation", handler.CalculationInstallment(deps.Services.installment))

		r.Route("/payment", func(r chi.Router) {
			r.Use(middleware.CheckAuthorization())

			r.Get("/calculation/{txn_id}", handler.CalculationTransactionInstallment(deps.Services.installment))
			r.Post("/", handler.PaymentInstallment(deps.Services.installment))
		})
	})
}
