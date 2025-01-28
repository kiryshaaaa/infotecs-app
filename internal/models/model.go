package models

type Wallet struct {
	ID      string `json:"walletid"`
	Balance int    `json:"balance"`
}

type Transaction struct {
	ID           int    `json:"transactionid"`
	FromWalletID string `json:"from"`
	ToWalletId   string `json:"to"`
	Amount       int    `json:"amount"`
}
