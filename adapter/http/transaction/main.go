package transaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"finance-project/util"
	"finance-project/models/transaction"
)

// GetTransactions returns all transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title: "Salário",
			Amount: 1288.0,
			Type: 0,
			CreatedAt: util.StringToTime("2024-10-13T11:01:05"),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(transactions)
}

// CreateATransaction creates a new transaction
func CreateATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &res)


	fmt.Println(res)

}