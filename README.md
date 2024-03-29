# CLI Tasks eFishery

Offline first app running in CLI written in GO using PouchDB / Apache CouchDB

## Objectives

Applications that work as well offline as they do online.

## Prerequisites

- Go
- PouchDB or CouchDB
- Kivik (PouchDB & CouchDB Library for GO)

## Geting Started

### Install GO

```
brew install go
```

### Install Kivik

```
go get -u github.com/go-kivik/kivik
go get -u github.com/go-kivik/couchdb
```

### Setup Environment

```
cp .env.example .env
```

### Run the CLI

```
go run app.go
```

## Running the test

```
go test cmd/command_test.go
```
