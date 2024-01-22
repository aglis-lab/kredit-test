package handler

import (
	"fmt"
	"kreditplus/src/app"
	"kreditplus/src/response"
	"kreditplus/src/v1/contract"
	"net/http"
)

func LoginCustomer(service CustomerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := app.Tracer().Start(r.Context(), "LoginCustomerHandler")
		defer span.End()

		req, err := contract.ValidateJSONRequest[contract.LoginCustomerRequest](ctx, r)
		if err != nil {
			span.RecordError(fmt.Errorf("ValidateLoginCustomerRequest, err: %v", err))
			response.JSONBadRequestResponse(ctx, w)
			return
		}

		resp, err := service.Login(ctx, &req)
		if err != nil {
			span.RecordError(fmt.Errorf("LoginCustomer, err: %v", err))
			response.JSONInternalServerError(ctx, w)
			return
		}

		response.JSONSuccessResponse(ctx, w, resp)
	}
}

func LimitCustomer(service CustomerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := app.Tracer().Start(r.Context(), "LimitCustomerHandler")
		defer span.End()

		resp, err := service.Limit(ctx)
		if err != nil {
			span.RecordError(fmt.Errorf("LimitCustomer, err: %v", err))
			response.JSONInternalServerError(ctx, w)
			return
		}

		response.JSONSuccessResponse(ctx, w, resp)
	}
}
