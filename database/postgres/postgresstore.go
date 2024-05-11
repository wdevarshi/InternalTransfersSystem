package postgres

import (
	"context"
	"database/sql"
	"github.com/wdevarshi/InternalTransfersSystem/database"
)

type Store struct {
	DB *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		DB: db,
	}
}

func (s *Store) CreateAccount(ctx context.Context, account *database.Account) error {
	//TODO implement me
	panic("implement me")
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
