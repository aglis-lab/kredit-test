package handler

import (
	"fmt"
	"kreditplus/src/app"
	"kreditplus/src/response"
	"kreditplus/src/v1/contract"
	"net/http"
)

func CalculationInstallment(svc InstallmentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := app.Tracer().Start(r.Context(), "CalculationInstallmentHandler")
		defer span.End()

		req, err := contract.ValidateJSONRequest[contract.CalculationInstallmentRequest](ctx, r)
		if err != nil {
			span.RecordError(fmt.Errorf("ValidateTransactionRequest, err: %v", err))
			response.JSONBadRequestResponse(ctx, w)
			return
		}

		resp, err := svc.Calculation(ctx, &req)
		if err != nil {
			span.RecordError(fmt.Errorf("CheckTransaction, err: %v", err))
			response.JSONInternalServerError(ctx, w)
			return
		}

		response.JSONSuccessResponse(ctx, w, resp)
	}
}

func CalculationTransactionInstallment(svc InstallmentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := app.Tracer().Start(r.Context(), "CalculationTransactionInstallmentHandler")
		defer span.End()

		txnID := contract.ValidateParamRequest(ctx, "txn_id")
		resp, err := svc.CalculationTransaction(ctx, txnID)
		if err != nil {
			span.RecordError(fmt.Errorf("CheckTransaction, err: %v", err))
			response.JSONInternalServerError(ctx, w)
			return
		}

		response.JSONSuccessResponse(ctx, w, resp)
	}
}

func PaymentInstallment(svc InstallmentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := app.Tracer().Start(r.Context(), "PaymentInstallmentHandler")
		defer span.End()

		req, err := contract.ValidateJSONRequest[contract.PaymentInstallmentRequest](ctx, r)
		if err != nil {
			span.RecordError(fmt.Errorf("ValidateTransactionRequest, err: %v", err))
			response.JSONBadRequestResponse(ctx, w)
			return
		}

		resp, err := svc.Payment(ctx, &req)
		if err != nil {
			span.RecordError(fmt.Errorf("CheckTransaction, err: %v", err))
			response.JSONInternalServerError(ctx, w)
			return
		}

		response.JSONSuccessResponse(ctx, w, resp)
	}
}
