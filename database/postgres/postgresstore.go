package postgres

import (
	"context"
	"database/sql"
	"github.com/wdevarshi/InternalTransfersSystem/database"
)

type Store struct {
	DB *sql.DB
}

func NewStore(db *sql.DB) database.InternalTransferSystemStore {
	return &Store{
		DB: db,
	}
}

func (s *Store) CreateAccount(ctx context.Context, account *database.Account) error {
	_, err := s.DB.Exec("INSERT INTO account (id, balance, time_created, last_modified, version) VALUES ($1, $2, $3, $4, $5)", account.ID, account.Balance, account.TimeCreated, account.LastModified, 0)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetAccount(ctx context.Context, accountID string) (*database.Account, error) {

	row := s.DB.QueryRow("SELECT id, balance, time_created, last_modified FROM account WHERE id = $1", accountID)

	account := &database.Account{}
	err := row.Scan(&account.ID, &account.Balance, &account.TimeCreated, &account.LastModified)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *Store) UpdateAccountWithTrx(ctx context.Context, account *database.Account) error {
	return nil
}

func (s *Store) AddTransaction(ctx context.Context, transaction *database.Transaction) error {

	_, err := s.DB.Exec("INSERT INTO trx_ledger (id, source_account_id, destination_account_id, amount, status, time_created, last_modified, version, error_reason) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", transaction.ID, transaction.SourceAccountID, transaction.DestinationAccountID, transaction.Amount, transaction.Status, transaction.TimeCreated, transaction.LastModified, transaction.Version, transaction.ErrorReason)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Close() error {
	return s.DB.Close()
}
