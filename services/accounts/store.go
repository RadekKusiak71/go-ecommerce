package accounts

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/RadekKusiak71/goEcom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetAccounts() ([]*types.Account, error) {
	rows, err := s.db.Query("SELECT * FROM accounts")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	accounts := make([]*types.Account, 0)
	for rows.Next() {
		acc, err := scanIntoAccount(rows)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		accounts = append(accounts, acc)
	}
	return accounts, nil
}

func (s *Store) GetAccountByID(accountID int) (*types.Account, error) {
	rows, err := s.db.Query("SELECT * FROM accounts WHERE id = $1", accountID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil, fmt.Errorf("user with id: %v wasn't found", accountID)
}

func (s *Store) GetAccountByEmail(email string) (*types.Account, error) {
	rows, err := s.db.Query("SELECT * FROM accounts WHERE email = $1", email)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil, fmt.Errorf("user with email: %v wasn't found", email)
}

func (s *Store) CreateAccount(account types.Account) (string, error) {
	_, err := s.db.Exec(`INSERT INTO accounts (first_name,last_name,email,password)
						VALUES ($1,$2,$3,$4)`, account.FirstName, account.LastName, account.Email, account.Password)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("error while creating account")
	}
	return "account created", nil
}

func scanIntoAccount(rows *sql.Rows) (*types.Account, error) {
	account := new(types.Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Email,
		&account.Password,
		&account.CreatedAt,
	)
	return account, err
}
