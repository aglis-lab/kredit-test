package handler

import (
	"context"
	"kreditplus/src/v1/contract"
)

type TransactionService interface {
	Create(context.Context, *contract.TransactionRequest) (contract.CreateTransactionResponse, error)
	Get(context.Context, string) (contract.GetTransactionResponse, error)
	Settlement(context.Context, *contract.SettlementTransactionRequest) (contract.GetTransactionResponse, error)
}

type InstallmentService interface {
	Calculation(context.Context, *contract.CalculationInstallmentRequest) (contract.CalculationInstallmentResponse, error)
	Payment(context.Context, *contract.PaymentInstallmentRequest) (contract.PaymentInstallmentResponse, error)
	CalculationTransaction(context.Context, string) (contract.CalculationInstallmentResponse, error)
}

type CustomerService interface {
	Login(context.Context, *contract.LoginCustomerRequest) (contract.LoginCustomerResponse, error)
	Limit(context.Context) (contract.LimitCustomerResponse, error)
}
