# IPFS WebApp

A simple web application written in Go to interact with IPFS.

## Running it

Start with installing dependencies:

```
go get -v ./...
```

And then run the program:

```
go run main.go
```

## Requests

Storing text values:

```
curl -X POST http://127.0.0.1:8081/upload_text --data "Hello World!"
```

Retrieving values:

```
curl -X GET http://127.0.0.1:8081/retrieve_text --data QmaC2xikV6UcK2LVc1TdTiDeH8s69N6URLGNbNwCB6FNQh
```