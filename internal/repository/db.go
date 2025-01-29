package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type Storage struct {
	db   *sql.DB
	psql squirrel.StatementBuilderType
}

func NewStorage() (*Storage, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	queryBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	storage := &Storage{
		db:   db,
		psql: queryBuilder,
	}

	if err := storage.CreateWalletsTable(); err != nil {
		return nil, fmt.Errorf("failed to create wallets table: %w", err)
	}
	if err := storage.InitializeWallets(); err != nil {
		return nil, fmt.Errorf("failed to initialize wallets: %w", err)
	}
	if err := storage.CreateTransactionsTable(); err != nil {
		return nil, fmt.Errorf("failed to create transactions table: %w", err)
	}

	return storage, nil
}
