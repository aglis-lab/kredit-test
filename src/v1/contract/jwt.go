package contract

import (
	"kreditplus/src/app"
	"kreditplus/src/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomerJWT struct {
	jwt.MapClaims
	ID        uint      `json:"id"`
	IssueDate time.Time `json:"issue_date"`
}

func GenerateJWT(customer entity.Customer) (string, error) {
	payload := CustomerJWT{ID: customer.ID, IssueDate: time.Now()}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return claims.SignedString([]byte(app.Config().SecretKey))
}
