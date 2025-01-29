package models

import "time"

type Transaction struct {
	ID         int       `db:"id"`
	FromWallet string    `db:"from_wallet"`
	ToWallet   string    `db:"to_wallet"`
	Amount     float64   `db:"amount"`
	Timestamp  time.Time `db:"timestamp"`
}
