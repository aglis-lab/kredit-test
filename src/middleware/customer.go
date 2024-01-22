package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"kreditplus/src/app"
	"kreditplus/src/v1/contract"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

const (
	AuthorizationKey = "Authorization"
)

type CustomerID struct{}

var CtxCustomerID = CustomerID{}

func JSONResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}

func ValidateAuthorization() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get(AuthorizationKey)
			if len(authorization) <= 7 {
				next.ServeHTTP(w, r)
				return
			}

			customerJwt := contract.CustomerJWT{}
			_, err := jwt.ParseWithClaims(authorization[7:], &customerJwt, func(t *jwt.Token) (interface{}, error) {
				return []byte(app.Config().SecretKey), nil
			})
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), CtxCustomerID, customerJwt.ID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func CheckAuthorization() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := GetAuthorizeCustomer(r.Context())
			if err != nil {
				JSONResponse(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func GetAuthorizeCustomer(ctx context.Context) (uint, error) {
	if v, ok := ctx.Value(CtxCustomerID).(uint); ok {
		return v, nil
	}

	return 0, errors.New("unauthorized")
}
