-- jdbc:postgresql://localhost:5432/internaltransferssystem
-- psql -d test_db -U postgres
-- docker stop $(docker ps -aq)
-- docker rm $(docker ps -aq)
-- docker rmi $(docker images -q)
-- docker volume rm $(docker volume ls -q)

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

-- 1. Remove the data and schema created during the testin
-- 2. Hook in the migration script to create the database and the tables.