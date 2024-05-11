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

func (s *Store) UpdateAccount(ctx context.Context, account *database.Account) error {
	_, err := s.DB.Exec("UPDATE account SET balance = $1, last_modified = $2 WHERE id = $3", account.Balance, account.LastModified, account.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) AddTransaction(ctx context.Context, transaction *database.Transaction) error {
	_, err := s.DB.Exec("INSERT INTO trx_ledger (id, source_account_id, destination_account_id, amount, status, time_created, last_modified) VALUES ($1, $2, $3, $4, $5, $6, $7)", transaction.ID, transaction.SourceAccountID, transaction.DestinationAccountID, transaction.Amount, transaction.Status, transaction.TimeCreated, transaction.LastModified)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetAllSourceTransactions(ctx context.Context, sourceAccountId string) ([]*database.Transaction, error) {

	rows, err := s.DB.Query("SELECT id, source_account_id, destination_account_id, amount, status, time_created, last_modified FROM trx_ledger WHERE source_account_id = $1", sourceAccountId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*database.Transaction
	for rows.Next() {
		transaction := &database.Transaction{}
		err := rows.Scan(&transaction.ID, &transaction.SourceAccountID, &transaction.DestinationAccountID, &transaction.Amount, &transaction.Status, &transaction.TimeCreated, &transaction.LastModified)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (s *Store) GetAllDestinationTransactions(ctx context.Context, destinationAccountId string) ([]*database.Transaction, error) {
	rows, err := s.DB.Query("SELECT id, source_account_id, destination_account_id, amount, status, time_created, last_modified FROM trx_ledger WHERE destination_account_id = $1", destinationAccountId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*database.Transaction
	for rows.Next() {
		transaction := &database.Transaction{}
		err := rows.Scan(&transaction.ID, &transaction.SourceAccountID, &transaction.DestinationAccountID, &transaction.Amount, &transaction.Status, &transaction.TimeCreated, &transaction.LastModified)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (s *Store) Close() error {
	return s.DB.Close()
}
