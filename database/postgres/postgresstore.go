package postgres

import (
	"context"
	"database/sql"
	"fmt"
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
	_, err := s.DB.ExecContext(ctx, "INSERT INTO account (id, balance, time_created, last_modified, version) VALUES ($1, $2, $3, $4, $5)", account.ID, account.Balance, account.TimeCreated, account.LastModified, 0)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetAccount(ctx context.Context, accountID string) (*database.Account, error) {

	row := s.DB.QueryRowContext(ctx, "SELECT * FROM account WHERE id = $1", accountID)

	account := &database.Account{}
	err := row.Scan(&account.ID, &account.Balance, &account.TimeCreated, &account.LastModified, &account.Version)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *Store) UpdateAccountWithTrx(ctx context.Context, sourceAccount *database.Account, destinationAccount *database.Account) error {
	trx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = trx.ExecContext(ctx, "UPDATE account SET balance = $1, last_modified = $2, version = $3 WHERE id = $4 and version = $5", sourceAccount.Balance, sourceAccount.LastModified, sourceAccount.Version+1, sourceAccount.ID, sourceAccount.Version)
	if err != nil {
		err := trx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	_, err = trx.ExecContext(ctx, "UPDATE account SET balance = $1, last_modified = $2, version = $3 WHERE id = $4 and version = $5", destinationAccount.Balance, destinationAccount.LastModified, destinationAccount.Version+1, destinationAccount.ID, destinationAccount.Version)
	fmt.Println(err)
	if err != nil {
		err := trx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	err = trx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) AddTransaction(ctx context.Context, transaction *database.Transaction) error {

	_, err := s.DB.Exec("INSERT INTO trx_ledger (id, source_account_id, destination_account_id, amount, status, time_created, last_modified, version, error_reason) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", transaction.ID, transaction.SourceAccountID, transaction.DestinationAccountID, transaction.Amount, transaction.Status, transaction.TimeCreated, transaction.LastModified, transaction.Version, transaction.ErrorReason)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateTransaction(ctx context.Context, transaction *database.Transaction) error {

	_, err := s.DB.Exec("UPDATE trx_ledger SET status = $1, last_modified = $2, version = $3, error_reason = $4 WHERE id = $5 and version = $6", transaction.Status, transaction.LastModified, transaction.Version+1, transaction.ErrorReason, transaction.ID, transaction.Version)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Close() error {
	return s.DB.Close()
}
