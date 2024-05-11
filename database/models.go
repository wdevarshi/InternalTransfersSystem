package database

import "time"

// CREATE TABLE "account"
// (
// "id"            uuid                     NOT NULL PRIMARY KEY,
// "balance"       DOUBLE PRECISION         NOT NULL,
// "time_created"  timestamp with time zone NOT NULL,
// "last_modified" timestamp with time zone NOT NULL,
// "version"       integer                  NOT NULL
// );
type Account struct {
	ID           string    `json:"id"`
	Balance      float64   `json:"balance"`
	TimeCreated  time.Time `json:"time_created"`
	LastModified time.Time `json:"last_modified"`
	Version      int       `json:"version"`
}

// create type trx_status as enum (
// 'init',
// 'success',
// 'failed'
// );
//
// CREATE TABLE "trx_ledger"
// (
// "id"                     uuid                     NOT NULL PRIMARY KEY,
// "source_account_id"      uuid                     NOT NULL,
// "destination_account_id" uuid                     NOT NULL,
// "amount"                 DOUBLE PRECISION         NOT NULL,
// "status"                 trx_status               NOT NULL,
// "time_created"           timestamp with time zone NOT NULL,
// "last_modified"          timestamp with time zone NOT NULL,
// "version"                integer                  NOT NULL,
// "error_reason"           text
// );
//
// CREATE INDEX "source_account_idx" on "trx_ledger" ("source_account_id");
// CREATE INDEX "destinatio_account_idnx" on "trx_ledger" ("destination_account_id");
type Transaction struct {
	ID                   string    `json:"id"`
	SourceAccountID      string    `json:"source_account_id"`
	DestinationAccountID string    `json:"destination_account_id"`
	Amount               float64   `json:"amount"`
	Status               string    `json:"status"`
	TimeCreated          time.Time `json:"time_created"`
	LastModified         time.Time `json:"last_modified"`
	Version              int       `json:"version"`
	ErrorReason          string    `json:"error_reason"`
}

const (
	TransactionStatusInit    = "init"
	TransactionStatusSuccess = "success"
	TransactionStatusFailed  = "failed"
)

const (
	ErrorReason_InsufficientBalance = "insufficient balance"
)

const (
	INIT_Version = 0
)
