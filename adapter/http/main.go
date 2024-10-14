package http

import (
	"finance-project/adapter/http/actuator"
	"finance-project/adapter/http/transaction"
	"fmt"
	"net/http"
	
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Init initializes the HTTP server
func Init(){
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
	fmt.Println("Listening on port 8080")
}