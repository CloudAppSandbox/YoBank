package main

import (
	"bankapp/accounts"
	"bankapp/accounts/api"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {

	initAppParams()

	initDependencies()
}

func main() {

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	StartHTTPServer()
	<-done

}

func StartHTTPServer() {

	r := mux.NewRouter()
	r.HandleFunc("/new", api.NewAccount).Methods("POST")
	r.HandleFunc("/close", api.CloseAccount).Methods("POST")
	r.HandleFunc("/deposit", api.Deposit).Methods("POST")
	r.HandleFunc("/withdraw", api.Withdraw).Methods("POST")
	r.HandleFunc("/transfer", api.Transfer).Methods("POST")
	r.HandleFunc("/", api.Get).Methods("GET")


	go func() {
		if err := http.ListenAndServe(":4900",r); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
		log.Print(fmt.Sprintf("Listening on address "))
	}()

	/*if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}*/
}

func initDependencies() {

	datasourceURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("database.host"), viper.GetInt("database.port"), viper.GetString("database.user"), viper.GetString("database.password"), viper.GetString("database.dbname"))

	db, err := sql.Open("postgres", datasourceURL)
	if err != nil {
		panic("Error opening connection to DB.")
	}
	repo := accounts.NewAccountsRepository(db)
	service := accounts.NewService(repo)
	api.New(service)
	log.Print(fmt.Sprintf("Finished setting up dependencies for app."))
}

func initAppParams() {
	viper.SetConfigName("application")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("properties")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	log.Print(fmt.Sprintf("Finished setting up viper from application.properties."))

}
