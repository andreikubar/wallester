# wallester test

## Initialize the database

### Connection parameters

Edit the database connection parameters in `db/connection.go`

### Run the data import procedure

```
cd db_init
go run initDb.go
```

### Execute unit tests

In the root directory:
```
go test
```

## Start the server

In the root directory:
```
go run wallester.go
```
