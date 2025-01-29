package services

import (
	"fmt"

	"github.com/kiryshaaaa/infotecs-app/internal/dto"
	"github.com/kiryshaaaa/infotecs-app/internal/repository"
)

type WalletService struct {
	repo *repository.Storage
}

func NewWalletService(repo *repository.Storage) *WalletService {
	return &WalletService{repo}
}

func (s *WalletService) GetBalance(address string) (*dto.WalletDTO, error) {
	balance, err := s.repo.GetBalance(address)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}
	return &dto.WalletDTO{
		Address: address,
		Balance: balance,
	}, nil
}
