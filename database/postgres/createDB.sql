-- jdbc:postgresql://localhost:5432/internaltransferssystem
-- psql postgres
CREATE
DATABASE internaltransferssystem;
-- \c internaltransferssystem
CREATE TABLE "account"
(
    "id"            uuid                     NOT NULL PRIMARY KEY,
    "balance"       DOUBLE PRECISION         NOT NULL,
    "time_created"  timestamp with time zone NOT NULL,
    "last_modified" timestamp with time zone NOT NULL,
    "version"       integer                  NOT NULL
);

create type trx_status as enum (
  'init',
  'success',
  'failed'
);

CREATE TABLE "trx_ledger"
(
    "id"                     uuid                     NOT NULL PRIMARY KEY,
    "source_account_id"      uuid                     NOT NULL,
    "destination_account_id" uuid                     NOT NULL,
    "amount"                 DOUBLE PRECISION         NOT NULL,
    "status"                 trx_status               NOT NULL,
    "time_created"           timestamp with time zone NOT NULL,
    "last_modified"          timestamp with time zone NOT NULL,
    "version"                integer                  NOT NULL,
    "error_reason"           text
);

CREATE INDEX "source_account_idx" on "trx_ledger" ("source_account_id");
CREATE INDEX "destination_account_idx" on "trx_ledger" ("destination_account_id");