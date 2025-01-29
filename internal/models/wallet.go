package models

type Wallet struct {
	ID      int     `db:"id"`
	Address string  `db:"address"`
	Balance float64 `db:"balance"`
}
