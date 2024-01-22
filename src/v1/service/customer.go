package service

import (
	"context"
	"kreditplus/src/app"
	"kreditplus/src/entity"
	"kreditplus/src/middleware"
	"kreditplus/src/tracer"
	"kreditplus/src/v1/contract"

	"golang.org/x/crypto/bcrypt"
)

func (service CustomerService) Login(ctx context.Context, req *contract.LoginCustomerRequest) (contract.LoginCustomerResponse, error) {
	ctx, span := app.Tracer().Start(ctx, "LoginCustomerService")
	defer span.End()

	resp := contract.LoginCustomerResponse{}

	customer := entity.Customer{}
	err := app.GormDB().WithContext(ctx).Where("email", req.Email).First(&customer).Error
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(req.Password)); err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	jwtStr, err := contract.GenerateJWT(customer)
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	resp.Token = jwtStr

	return resp, nil
}

func (svc CustomerService) Limit(ctx context.Context) (contract.LimitCustomerResponse, error) {
	ctx, span := app.Tracer().Start(ctx, "LimitCustomerService")
	defer span.End()

	resp := contract.LimitCustomerResponse{}

	db := app.GormDB().WithContext(ctx)

	// Get Customer Data
	customerID, err := middleware.GetAuthorizeCustomer(ctx)
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	customer := entity.Customer{}
	err = db.Preload("Limits").First(&customer, customerID).Error
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	for _, item := range customer.Limits {
		resp.Limits = append(resp.Limits, contract.LimitCustomerItem{
			Limit:     item.Limit,
			Period:    item.Period,
			UsedLimit: item.GetUsedLimit(),
		})
	}

	resp.Interest = customer.Interest
	resp.Risk = customer.Risk

	return resp, nil
}
