package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCrad) error
	GetCreditCard(creditCard CreditCrad) (CreditCrad, error)
	CreateCreditCard(creditCard CreditCrad) error
}

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	Description  string
	Store        string
	CreditCradId string
	CrteatedAt   time.Time
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.ID = uuid.NewV4().String()
	t.CrteatedAt = time.Now()
	return t
}

func (t *Transaction) ProcessAndValidate(creditCrad *CreditCrad) {
	if t.Amount+creditCrad.Balance > creditCrad.Limit {
		t.Status = "rejected"
	} else {
		t.Status = "approved"
		creditCrad.Balance = creditCrad.Balance + t.Amount
	}
}
