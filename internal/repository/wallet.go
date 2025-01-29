package repository

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/Masterminds/squirrel"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

func GenerateWalletAddress() string {
	b := make([]byte, 64)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (s *Storage) CreateWalletsTable() error {
	query := `
    CREATE TABLE IF NOT EXISTS wallets (
		id SERIAL PRIMARY KEY,
        address CHAR(64) UNIQUE,
        balance DECIMAL(10, 2) NOT NULL CHECK (balance >= 0)
    );`

	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create wallets table: %w", err)
	}

	log.Println("Wallets table created or already exists!")
	return nil
}

func (s *Storage) InitializeWallets() error {
	var count int
	query, args, err := s.psql.
		Select("COUNT(*)").
		From("wallets").
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	err = s.db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check wallets: %w", err)
	}

	if count > 0 {
		log.Println("Wallets already exist, skipping initialization.")
		return nil
	}

	for i := 0; i < 10; i++ {
		address := GenerateWalletAddress()
		query, args, err := s.psql.
			Insert("wallets").
			Columns("address", "balance").
			Values(address, 100.0).
			ToSql()
		if err != nil {
			return fmt.Errorf("failed to build insert query: %w", err)
		}

		_, err = s.db.Exec(query, args...)
		if err != nil {
			return fmt.Errorf("failed to insert wallet: %w", err)
		}
	}

	log.Println("10 wallets initialized with 100.0 balance each!")
	return nil
}

func (s *Storage) GetBalance(address string) (float64, error) {
	var balance float64
	query, args, err := s.psql.
		Select("balance").
		From("wallets").
		Where(squirrel.Eq{"address": address}).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build select query: %w", err)
	}

	err = s.db.QueryRow(query, args...).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("wallet not found")
		}
		return 0, fmt.Errorf("failed to get balance: %w", err)
	}
	return balance, nil
}
