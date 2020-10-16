package api

import "bankapp/accounts"

type Account struct {
	Type    string  `validator:"oneof=savings current"`
	Name    string  `validator:"gte=10"`
	PAN     string  `validator:"len=10"`
	Balance float64 `validator:"gte=0"`
}

type AuthInfo struct {
	AccountNumber int32 `validator:"len=15"`
	PassCode      string
}

type AccountService interface {
	CreateAccount(acc accounts.Account) accounts.AuthInfo
	CloseAccount(accountNumber int32) float64
	MakeDeposit(accountNumber int32, auth accounts.AuthInfo, amount float64) float64
	Withdraw(accountNumber int32, auth accounts.AuthInfo, amount float64) float64
	Transfer(accountNumber int32, auth accounts.AuthInfo, amount float64) bool
}
