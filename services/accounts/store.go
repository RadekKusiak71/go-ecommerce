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

func (s *Store) UpdateAccount(account types.Account, accountID int) (*types.Account, error) {
	_, err := s.db.Query(`UPDATE accounts SET first_name = $1 , last_name = $2 , email = $3 WHERE id = $4`,
		account.FirstName,
		account.LastName,
		account.Email,
		accountID,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return s.GetAccountByID(accountID)
}

func (s *Store) DeleteAccount(accountID int) (string, error) {
	res, err := s.db.Exec("DELETE FROM accounts WHERE id = $1", accountID)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if count, err := res.RowsAffected(); err != nil || count == 0 {
		return "account don't exists", err
	}
	return "account deleted", nil
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
	_, err := s.db.Exec(`INSERT INTO accounts (first_name,last_name,email,password_hash)
						VALUES ($1,$2,$3,$4)`, account.FirstName, account.LastName, account.Email, account.Password)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("account with this email already exists")
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
