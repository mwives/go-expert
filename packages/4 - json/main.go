package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  string  `json:"account_number"`
	Balance float64 `json:"balance"`
}

func main() {
	account := Account{
		Number:  "123",
		Balance: 100.0,
	}

	// Convert struct to JSON
	res, err := json.Marshal(account)
	if err != nil {
		println(err)
	}
	println(string(res))

	// Convert struct to JSON and print to stdout
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		println(err)
	}

	// Convert JSON to struct
	rawJson := []byte(`{"account_number":"234","balance":200.0}`)
	var account2 Account
	err = json.Unmarshal(rawJson, &account2)
	if err != nil {
		println(err)
	}
	println(account2.Number)
}
