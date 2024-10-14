package http

import (
	"net/http"
	"finance-project/adapter/http/transaction"
	"finance-project/adapter/http/actuator"
	"fmt"
)

// Init initializes the HTTP server
func Init(){
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
	fmt.Println("Listening on port 8080")
}