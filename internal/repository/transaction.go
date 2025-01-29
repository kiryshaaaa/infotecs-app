package repository

import (
	"fmt"
	"log"

	"github.com/kiryshaaaa/infotecs-app/internal/models"
)

func (s *Storage) CreateTransactionsTable() error {
	query := `
    CREATE TABLE IF NOT EXISTS transactions (
        id SERIAL PRIMARY KEY,
        from_wallet CHAR(64) NOT NULL,
        to_wallet CHAR(64) NOT NULL,
        amount DECIMAL(10, 2) NOT NULL CHECK (amount > 0),
        timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create transactions table: %w", err)
	}
	log.Println("Transactions table created or already exists!")
	return nil
}

func (s *Storage) InsertTransaction(from, to string, amount float64) error {
	query, args, err := s.psql.
		Insert("transactions").
		Columns("from_wallet", "to_wallet", "amount").
		Values(from, to, amount).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build insert query: %w", err)
	}

	_, err = s.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %w", err)
	}

	return nil
}

func (s *Storage) GetLastNTransactions(n int) ([]models.Transaction, error) {
	query, args, err := s.psql.
		Select("id", "from_wallet", "to_wallet", "amount", "timestamp").
		From("transactions").
		OrderBy("timestamp DESC").
		Limit(uint64(n)).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute select query: %w", err)
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var tx models.Transaction

		if err := rows.Scan(&tx.ID, &tx.FromWallet, &tx.ToWallet, &tx.Amount, &tx.Timestamp); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		transactions = append(transactions, tx)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return transactions, nil
}
