-- jdbc:postgresql://localhost:5432/internaltransferssystem
-- psql postgres
CREATE
DATABASE internaltransferssystem;
-- \c internaltransferssystem
CREATE TABLE "account"
(
    "id"            uuid                     NOT NULL PRIMARY KEY,
    "balance"       varchar(50)              NOT NULL,
    "time_created"  timestamp with time zone NOT NULL,
    "last_modified" timestamp with time zone NOT NULL
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
    "destinatio_account_idn" uuid                     NOT NULL,
    "amount"                 DOUBLE PRECISION         NOT NULL,
    "status"                 trx_status               NOT NULL,
    "time_created"           timestamp with time zone NOT NULL,
    "last_modified"          timestamp with time zone NOT NULL
);

CREATE INDEX "source_account_idx" on "trx_ledger" ("source_account_id");
CREATE INDEX "destinatio_account_idnx" on "trx_ledger" ("destinatio_account_idn");