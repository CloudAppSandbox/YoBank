package api

import (
	"bankapp/accounts"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type requestHandler struct {
	service AccountService
}

var (
	handlerInstance requestHandler
	handlerInit     sync.Once
)

func New(accService AccountService) {
	handlerInit.Do(func() {
		handlerInstance = requestHandler{service: accService}
	})
}

func HandlerInstance() *requestHandler {
	return &handlerInstance
}

func NewAccount(writer http.ResponseWriter, request *http.Request) {
	log.Print(fmt.Print("Parsing Request for Create Account."))

	body := request.Body

	bytes, requestReadErr := ioutil.ReadAll(body)

	if requestReadErr != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Unexpected Message Parsed."))
	}

	var acc accounts.Account
	marshalErr := json.Unmarshal(bytes, &acc)
	if marshalErr != nil {
		//TODO handle possible validation issues
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Unexpected Message Parsed."))
	}


	log.Print("Passing to service from handler.")
	auth := HandlerInstance().service.CreateAccount(acc)

	resp, err := json.Marshal(&auth)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to create, try again."))
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(resp))

}

func CloseAccount(writer http.ResponseWriter, request *http.Request) {

}
func Deposit(writer http.ResponseWriter, request *http.Request) {

}
func Withdraw(writer http.ResponseWriter, request *http.Request) {

}
func Transfer(writer http.ResponseWriter, request *http.Request) {

}

func Get(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Welcome"))
}
