package v1

import (
	"context"
	"kreditplus/src/v1/service"
)

type repositories struct {
}

type services struct {
	transaction service.TransactionService
	installment service.InstallmentService
	customer    service.CustomerService
}

type Dependency struct {
	Repositories *repositories
	Services     *services
}

func Dependencies(ctx context.Context) *Dependency {
	repositories := repositories{}

	services := services{}

	return &Dependency{
		Repositories: &repositories,
		Services:     &services,
	}
}
