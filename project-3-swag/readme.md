# Swag project

## DB Migration

To run db migration:
1. Install migrate CLI

```sh
brew install golang-migrate
```

2. Run migration

```sh
migrate -source file:migration_path -database "postgres://user:password@localhost:5432/db_name sslmode=disable" up 1
```

Example:

```sh
migrate -source file:./db/migrations/ -database "postgres://user:password@localhost:5432/db_name sslmode=disable" up 1
```

## Run test with different db

Run in terminal


```sh
PSQL_URL=postgres://username@127.0.0.1:5432/swag_test?sslmode=disable go test -v
```