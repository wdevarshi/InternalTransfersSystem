package database

import (
	"context"
)

type InternalTransferSystemStore interface {
	CreateAccount(ctx context.Context, account *Account) error
	GetAccount(ctx context.Context, accountID string) (*Account, error)
	UpdateAccountWithTrx(ctx context.Context, sourceAccount *Account, destinationAccount *Account) error
	AddTransaction(ctx context.Context, transaction *Transaction) error
	UpdateTransaction(ctx context.Context, transaction *Transaction) error
	Close() error
}
