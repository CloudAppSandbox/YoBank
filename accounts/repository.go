package accounts

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

var (
	repoInstance Repository
	repoInit     sync.Once
)

type Repository struct {
	datastore *sql.DB
}

func NewAccountsRepository(dbInstance *sql.DB) *Repository {
	repoInit.Do(func() {
		repoInstance = Repository{datastore: dbInstance}
	})
	return &repoInstance
}

func (a Repository) CreateAccount(acc Account) AuthInfo {
	log.Print("Calling create in DB.")
	log.Print(fmt.Sprintf("DB details: %s", a.datastore.Stats()))
	return AuthInfo{}
}

func (a Repository) CloseAccount(accountNumber int32) float64 {
	panic("implement me")
}

func (a Repository) MakeDeposit(accountNumber int32, auth AuthInfo, amount float64) float64 {
	panic("implement me")
}

func (a Repository) Withdraw(accountNumber int32, auth AuthInfo, amount float64) float64 {
	panic("implement me")
}

func (a Repository) Transfer(accountNumber int32, auth AuthInfo, amount float64) bool {
	panic("implement me")
}
