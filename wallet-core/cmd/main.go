package main

import (
	"database/sql"
	"fmt"
	"github.com/IamP5/ms-wallet/wallet-core/internal/database"
	"github.com/IamP5/ms-wallet/wallet-core/internal/event"
	"github.com/IamP5/ms-wallet/wallet-core/internal/usecase/create_account"
	"github.com/IamP5/ms-wallet/wallet-core/internal/usecase/create_client"
	"github.com/IamP5/ms-wallet/wallet-core/internal/usecase/create_transaction"
	"github.com/IamP5/ms-wallet/wallet-core/internal/web"
	"github.com/IamP5/ms-wallet/wallet-core/internal/web/webserver"
	"github.com/IamP5/ms-wallet/wallet-core/pkg/events"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDipsatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	//eventDipsatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDipsatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTranscation)

	webserver.Start()
}
