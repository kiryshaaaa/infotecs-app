package services

import (
	"fmt"
	"log"
	"time"

	"github.com/kiryshaaaa/infotecs-app/internal/dto"
	"github.com/kiryshaaaa/infotecs-app/internal/repository"
)

type TransactionService struct {
	repo *repository.Storage
}

func NewTransactionService(repo *repository.Storage) *TransactionService {
	return &TransactionService{repo}
}

func (s *TransactionService) GetLastN(n int) ([]dto.TransactionDTO, error) {
	transactions, err := s.repo.GetLastNTransactions(n)
	if err != nil {
		return nil, fmt.Errorf("failed to get last N transactions: %w", err)
	}

	var result []dto.TransactionDTO
	for _, tx := range transactions {
		transaction := dto.TransactionDTO{
			ID:         tx.ID,
			FromWallet: tx.FromWallet,
			ToWallet:   tx.ToWallet,
			Amount:     tx.Amount,
			Timestamp:  tx.Timestamp.Format(time.RFC3339),
		}
		result = append(result, transaction)
	}

	return result, nil
}

func (s *TransactionService) Send(from, to string, amount float64) error {
	if from == to {
		return fmt.Errorf("sender and recipient cannot be the same")
	}

	senderBalance, err := s.repo.GetBalance(from)
	if err != nil {
		return fmt.Errorf("failed to get sender balance: %w", err)
	}

	if senderBalance < amount {
		return fmt.Errorf("insufficient funds")
	}

	err = s.repo.UpdateBalance(from, -amount)
	if err != nil {
		return fmt.Errorf("failed to update sender balance: %w", err)
	}

	err = s.repo.UpdateBalance(to, amount)
	if err != nil {
		_ = s.repo.UpdateBalance(from, amount)
		return fmt.Errorf("failed to update recipient balance: %w", err)
	}

	err = s.repo.InsertTransaction(from, to, amount)
	if err != nil {
		_ = s.repo.UpdateBalance(from, amount)
		_ = s.repo.UpdateBalance(to, -amount)
		return fmt.Errorf("failed to insert transaction: %w", err)
	}

	log.Printf("Transferred %.2f from %s to %s", amount, from, to)
	return nil
}
