package accounts

import (
	"log"
	"sync"
)

var (
	transactionInstance transactionHandler
	handlerInit         sync.Once
)

type holder interface {
	CreateAccount(acc Account) AuthInfo
	CloseAccount(accountNumber int32) float64
	MakeDeposit(accountNumber int32, auth AuthInfo, amount float64) float64
	Withdraw(accountNumber int32, auth AuthInfo, amount float64) float64
	Transfer(accountNumber int32, auth AuthInfo, amount float64) bool
}

type transactionHandler struct {
	accountHolder holder
}

func NewService(holderInstance holder) *transactionHandler {
	handlerInit.Do(func() {
		transactionInstance = transactionHandler{accountHolder: holderInstance}
	})
	return &transactionInstance
}

func (t transactionHandler) CreateAccount(acc Account) AuthInfo {
	log.Print("Passing to repository from service.")
	return t.accountHolder.CreateAccount(acc)
}

func (t transactionHandler) CloseAccount(accountNumber int32) float64 {
	return t.accountHolder.CloseAccount(accountNumber)
}

func (t transactionHandler) MakeDeposit(accountNumber int32, auth AuthInfo, amount float64) float64 {
	return t.accountHolder.MakeDeposit(accountNumber, auth, amount)
}

func (t transactionHandler) Withdraw(accountNumber int32, auth AuthInfo, amount float64) float64 {
	return t.accountHolder.Withdraw(accountNumber, auth, amount)
}

func (t transactionHandler) Transfer(accountNumber int32, auth AuthInfo, amount float64) bool {
	return t.accountHolder.Transfer(accountNumber, auth, amount)
}
