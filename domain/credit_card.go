package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreditCrad struct {
	ID              string
	Name            string
	Number          string
	ExpirationMonth int32
	ExpirationYear  int32
	CVV             int32
	Balance         float64
	Limit           float64
	CrteatedAt      time.Time
}

func NewCreditCard() *CreditCrad {
	c := &CreditCrad{}
	c.ID = uuid.NewV4().String()
	c.CrteatedAt = time.Now()
	return c
}
