Install migration tool

```
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Create database

```
CREATE DATABASE shortener;
```

Create migration

```
migrate create -ext sql -dir ./internal/store/sqlstore/mysql-migration create_table_users
```

Up

```
migrate -path ./internal/store/sqlstore/mysql-migration -database "mysql://root:pass@tcp(localhost:3306)/shortener" up
```

Down

```
migrate -path ./internal/store/sqlstore/mysql-migration -database "mysql://root:pass@tcp(localhost:3306)/shortener" down
```
