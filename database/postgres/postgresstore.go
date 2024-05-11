package postgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
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
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	_, err = s.DB.Exec("INSERT INTO account (id, balance, time_created, last_modified) VALUES ($1, $2, $3, $4)", uuid, account.Balance, account.TimeCreated, account.LastModified)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetAccount(ctx context.Context, accountID string) (*database.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) UpdateAccount(ctx context.Context, account *database.Account) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) AddTransaction(ctx context.Context, transaction *database.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) GetAllSourceTransactions(ctx context.Context, sourceAccountId string) ([]*database.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) GetAllDestinationTransactions(ctx context.Context, destinationAccountId string) ([]*database.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) Close() error {
	//TODO implement me
	panic("implement me")
}
