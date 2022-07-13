# Ektimo API

Ektimo API is the server-side application for Ektimo, a talent hiring system.

## Prerequisites

- [Go 1.18](https://golang.org/)

## Running locally

### Set up dependencies

```sh
make compose-up
```

### Run Ektimo API server

```sh
make run
```

Server will be live now on http://localhost:8080  
Swagger API documentation will be live on http://localhost:8080/swagger/index.html

### Tear down dependencies

```sh
make compose-down
```

## Generate ent code

```sh
make ent
```

Click [here](https://entgo.io/docs/getting-started/) to find more about ent

## Generate Swagger docs

```sh
make docs
```

## Run tests

```sh
make test
```