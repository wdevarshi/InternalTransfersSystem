package database

import "context"

type Store interface {
	CreateAccount(ctx context.Context, account *Account) error
	GetAccount(ctx context.Context, accountID string) (*Account, error)
	UpdateAccount(ctx context.Context, account *Account) error
	AddTransaction(ctx context.Context, transaction *Transaction) error
	GetAllSourceTransactions(ctx context.Context, sourceAccountId string) ([]*Transaction, error)
	GetAllDestinationTransactions(ctx context.Context, destinationAccountId string) ([]*Transaction, error)
	Close() error
}

//CREATE TABLE "account"
//(
//"id"            uuid                     NOT NULL PRIMARY KEY,
//"balance"       varchar(50)              NOT NULL,
//"time_created"  timestamp with time zone NOT NULL,
//"last_modified" timestamp with time zone NOT NULL
//);

type Account struct {
	ID           string `json:"id"`
	Balance      string `json:"balance"`
	TimeCreated  string `json:"time_created"`
	LastModified string `json:"last_modified"`
}

// create type trx_status as enum (
// 'init',
// 'success',
// 'failed'
// );
// CREATE TABLE "trx_ledger"
// (
// "id"                     uuid                     NOT NULL PRIMARY KEY,
// "source_account_id"      uuid                     NOT NULL,
// "destinatio_account_idn" uuid                     NOT NULL,
// "amount"                 DOUBLE PRECISION         NOT NULL,
// "status"                 trx_status               NOT NULL,
// "time_created"           timestamp with time zone NOT NULL,
// "last_modified"          timestamp with time zone NOT NULL
// );
type Transaction struct {
	ID                   string  `json:"id"`
	SourceAccountID      string  `json:"source_account_id"`
	DestinationAccountID string  `json:"destination_account_id"`
	Amount               float64 `json:"amount"`
	Status               string  `json:"status"`
	TimeCreated          string  `json:"time_created"`
	LastModified         string  `json:"last_modified"`
}