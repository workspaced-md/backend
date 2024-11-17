package db

import (
	"fmt"

	"github.com/arnavsurve/md/pkg/shared"
)

func (s *Store) CreateAccount(account *shared.Account) error {
    _, err := s.DB.Exec(`INSERT INTO accounts (email, username, password) VALUES ($1, $2, $3)`,
        account.Email, account.Username, account.Password)
    if err != nil {
        return fmt.Errorf("error creating account: %v", err)
    }
    return nil
}

func (s *Store) EditAccount(account *shared.Account) error {
    _, err := s.DB.Exec(`UPDATE accounts SET username = $1, email = $2 WHERE id = $3`,
        account.Email, account.Username, account.Password)
    if err != nil {
        return fmt.Errorf("error updating account: %v", err)
    }
    return nil
}

func (s *Store) GetAccountById(id int) (*shared.Account, error) {
	account := &shared.Account{}
	err := s.DB.QueryRow(`SELECT id, email, username FROM accounts WHERE id = $1`, id).Scan(
		&account.ID, &account.Email, &account.Username)
	if err != nil {
		return nil, fmt.Errorf("error getting account: %v", err)
	}
	return account, nil
}

func (s *Store) GetAccountByEmail(email string) (*shared.Account, error) {
	account := &shared.Account{}
	err := s.DB.QueryRow(`SELECT id, email, username FROM accounts WHERE email = $1`, email).Scan(
		&account.ID, &account.Email, &account.Username)
	if err != nil {
		return nil, fmt.Errorf("error getting account: %v", err)
	}
	return account, nil
}