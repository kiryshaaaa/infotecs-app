package dto

type TransactionDTO struct {
	ID         int     `json:"id"`
	FromWallet string  `json:"from"`
	ToWallet   string  `json:"to"`
	Amount     float64 `json:"amount"`
	Timestamp  string  `json:"timestamp"`
}

type SendRequestDTO struct {
	FromWallet string  `json:"from"`
	ToWallet   string  `json:"to"`
	Amount     float64 `json:"amount"`
}
