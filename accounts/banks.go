package accounts

type Account struct {
	AccountID int32
	Type      string
	Name      string
	PAN       string
	Balance   float64
	Status    string
	AuthInfo
}

type AuthInfo struct {
	AccountNumber int32
	PassCode      string
}
