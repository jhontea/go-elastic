# Project Elasticsearch Learn

Learning using elasticsearch on go languange programming.

## Getting Started

This project is an API for Create, Read, Update, and Delete in elasticsearch using `olivere` elasticsearch package.

### Installing

1. Use govendor for package management
```
go get -u github.com/kardianos/govendor
```
2. Install all package
```
govendor get -v
govendor sync
```
3. Generate swagger documentation
```
swag init
```
4. Running
```
go run -v main.go
```

## Run using refresh
1. Installation

```
go get github.com/markbates/refresh
```
2. First you'll want to create a `refresh.yml` configuration file:

```
refresh init
```
3. Once you have your configuration all set up, all you need to do is run it:

```
refresh run
```

## Swagger documentation
1. Download swag by using:
```sh
go get -u github.com/swaggo/swag/cmd/swag
```
2. Run `swag init` in the project's root folder which contains the `main.go` file. This will parse your comments and generate the required files (`docs` folder and `docs/docs.go`).
```sh
swag init
```

## Testing
```
go test ./... -v -cover
```