package handler

import (
	"fmt"
	"kreditplus/src/app"
	"kreditplus/src/response"
	"kreditplus/src/tracer"
	"kreditplus/src/v1/contract"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateTransaction(svc TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := app.Tracer().Start(r.Context(), "CreateTransactionHandler")
		defer span.End()

		req, err := contract.ValidateJSONRequest[contract.TransactionRequest](ctx, r)
		if err != nil {
			tracer.RecordError(span, fmt.Errorf("ValidateTransactionRequest, err: %v", err))
			response.JSONBadRequestResponse(ctx, w)
			return
		}

		resp, err := svc.Create(ctx, &req)
		if err != nil {
			tracer.RecordError(span, fmt.Errorf("CreateTransaction, err: %v", err))
			response.JSONInternalServerError(ctx, w)
			return
		}

		response.JSONSuccessResponse(ctx, w, resp)
	}
}

func GetTransaction(svc TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := app.Tracer().Start(r.Context(), "GetTransactionHandler")
		defer span.End()

		req := chi.URLParamFromCtx(ctx, "txn_id")

		resp, err := svc.Get(ctx, req)
		if err != nil {
			tracer.RecordError(span, fmt.Errorf("GetTransaction, err: %v", err))
			response.JSONInternalServerError(ctx, w)
			return
		}

		response.JSONSuccessResponse(ctx, w, resp)
	}
}

func SettlementTransaction(svc TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := app.Tracer().Start(r.Context(), "SettlementTransactionHandler")
		defer span.End()

		req, err := contract.ValidateJSONRequest[contract.SettlementTransactionRequest](ctx, r)
		if err != nil {
			tracer.RecordError(span, fmt.Errorf("ValidateTransactionRequest, err: %v", err))
			response.JSONBadRequestResponse(ctx, w)
			return
		}

		resp, err := svc.Settlement(ctx, &req)
		if err != nil {
			tracer.RecordError(span, fmt.Errorf("GetTransaction, err: %v", err))
			response.JSONInternalServerError(ctx, w)
			return
		}

		response.JSONSuccessResponse(ctx, w, resp)
	}
}
