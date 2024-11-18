package db

import (
	"fmt"

	"github.com/arnavsurve/workspaced/pkg/shared"
)

func (s *Store) CreateAccount(account *shared.Account) error {
	result := s.DB.Create(account)
	if result.Error != nil {
		return fmt.Errorf("error creating account: %v", result.Error)
	}
	return nil
}

func (s *Store) EditAccount(account *shared.Account) error {
	result := s.DB.Model(&shared.Account{}).Where("id = ?", account.ID).Updates(account)
	if result.Error != nil {
		return fmt.Errorf("error updating account: %v", result.Error)
	}
	return nil
}

func (s *Store) GetAccountById(id int) (*shared.Account, error) {
	account := &shared.Account{}
	result := s.DB.First(account, id)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting account: %v", result.Error)
	}
	return account, nil
}

func (s *Store) GetAccountByEmail(email string) (*shared.Account, error) {
	account := &shared.Account{}
	result := s.DB.Where("email = ?", email).First(account)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting account: %v", result.Error)
	}
	return account, nil
}

