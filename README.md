# InternalTransfersSystem

A project to learn concurreny and trx in postgres

## Getting started

This project requires Go to be installed.

On OS X with Homebrew you can just run `brew install go`.

On Linux use your package manager to install go or use the [go documentation](https://go.dev/doc/install)

Running it then should be as simple as:

```console
$ make run
```

## Makefile

The Makefile contains a number of useful commands to help you get started. Here are some of the most useful ones:
- `make help` - Prints the help
- `make test` - Runs the tests
- `make test-coverage` - Runs the tests and generates a coverage report
- `make test-coverage-html` - Runs the tests and generates a coverage report in HTML
- `make bench` - Runs the benchmarks
- `make lint` - Runs the linter
- `make run` - Runs the application
- `make build` - Builds the application
- `make generate` - Generates the code
- `make docker-build` - Builds the Docker image
- `make docker-run` - Runs the Docker image
- `make deps` - Installs the dependencies
- `make clean` - Cleans up the project

## Docker

This project also contains a Dockerfile to help you get started with Docker. To build the image, run:

```console
$ make docker-build
```

To run the image, run:

```console
$ make docker-run
```
## How to create DB
This project uses postgresql to store data. For the purpose this demo, create install postgresql and create a database called internalTransfersSystem. Moreover, for production, we prefer to run the postgres in docker with a volume mount to persist data. You can use the following command to create the database

```console
$ docker pull postgres
$ docker run --name internalTransfersSystem -e POSTGRES_PASSWORD=<> -e POSTGRES_USER=internalTransferSystemUser  -p 5432:5432 -d postgres -v pgdata:/var/lib/postgresql/data
```
Once you have the DB created, you need to create the tables. You can use init.sql file to create the tables. You can use the following command to create the tables

```console
-- This table is primarily used to serve the read requests. The goal of this table is to store the account details.
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

-- The goal of this table is to store the transaction details in a ledger. All the transactions will be stored in this table even for the failed transactions.
CREATE TABLE "trx_ledger"
(
    "id"                     uuid                     NOT NULL PRIMARY KEY,
    "source_account_id"      uuid                     NOT NULL,
    "destinatio_account_idn" uuid                     NOT NULL,
    "amount"                 DOUBLE PRECISION         NOT NULL,
    "status"                 trx_status               NOT NULL,
    "time_created"           timestamp with time zone NOT NULL,
    "last_modified"          timestamp with time zone NOT NULL,
    "version"                integer                  NOT NULL,
    "error_reason"           text
);

CREATE INDEX "source_account_idx" on "trx_ledger" ("source_account_id");
CREATE INDEX "destinatio_account_idnx" on "trx_ledger" ("destinatio_account_idn");
```

## Adding a new endpoint to the API

Our service is grpc first. We use [grpc-gateway] to automatically map HTTP requests to gRPC requests. This means that you can add a new endpoint to the API by adding a new rpc to `service InternalTransfersSystem` in `proto/internaltransferssystem.proto` file. Then, you can run `make generate` to generate grpc/http endpoints.

The file `serice/service.go` contains the implementation of the API and serves as the emtrypoint for the app. You can add your business logic there or any other package.

### Current endpoints - CURL

API to create account
```console
curl --location --request POST 'localhost:9091/api/v1/account/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_id" : "12359cc6-6ea6-41a6-9c52-79734acbf191",
    "initial_balance" : 100.0
}'
```
 
API to get account
```console
curl --location --request POST 'localhost:9091/api/v1/account/get' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_id" : "12359cc6-6ea6-41a6-9c52-79734aabf191"

}'
```

API to create transaction
```console
curl --location --request POST 'localhost:9091/api/v1/transaction/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "from_account_id" : "12359cc6-6ea6-41a6-9c52-79734aabf191",
    "to_account_id" : "12359fcc6-6ea6-41a6-9c52-79734acbf192",
    "amount" : 20.0

}'
```


### HTTP to gRPC mapping

We use [grpc-gateway] to automatically map HTTP requests to gRPC requests. You can find the mapping in the `proto/myapp.proto` file. This server is generated according to [custom options](https://cloud.google.com/service-infrastructure/docs/service-management/reference/rpc/google.api#http) in your gRPC definition.  You can find more information about the mapping [here](https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/)

### State machine for the Transaction

```console
    INIT --> SUCCESS
    INIT --> FAILED
```

### Future work

1. Add a worker pool to reconcile the transactions not in terminal state
2. Rate limiter, circuit breaker to prevent the abuse of the system
3. Add a cache to cache the account details like image URL, name, etc
4. Create load balancer and service discovery to balance the load 
5. Add a secret manager to manage the secrets 
6. A security layer to secure the system 
7. Add a CI/CD pipeline to automate the deployment 
8. DB slave to backup the data periodically for audit and disaster recovery 
9. Add a check for code quality, code smell, bugs, vulnerabilities, etc 
10. Add a code performance, stress, benchmark, profiling tests 
11. Include code best practices, Anti-pattern, review guidelines to follow in the readme

## Application configuration

This project uses [envconfig] to manage configuration as environment variables. You can find the configuration struct in `config/config.go`. You can also find the default values in the `config/config.go` file.

### Environment variables for local development

You can find the environment variables for local development in the `local.env` file. This file is used when you run `make run` or `make docker-run`.

### Coldbrew configuration options

A Large number of configuration options are prowered by [Coldbrew] and used as environment variables. You can find the list of environment variables [here](https://pkg.go.dev/github.com/go-coldbrew/core/config#Config).

## Logging

This project uses `go-coldbrew/log` to manage logging. You can find documentation [here](https://pkg.go.dev/github.com/go-coldbrew/log).

---
[envconfig]: https://github.com/kelseyhightower/envconfig
[grpc-gateway]: https://grpc-ecosystem.github.io/grpc-gateway/
[Coldbrew]: https://docs.coldbrew.cloud
