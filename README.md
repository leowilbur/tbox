# TBOX Backend Project [![CircleCI](https://circleci.com/gh/leowilbur/tbox/tree/master.svg?style=svg&circle-token=6c77be0831af8df121c16787d65045bea3285d64)](https://circleci.com/gh/leowilbur/tbox/tree/master)

This repository contains an implementation of the TBOX API. All the
endpoints have been described in the [Swagger definition](./swagger.yaml).

## Project structure

 - `.circleci` contains all files related to the CI process,
 - `migrations` contains all the SQL migrations,
 - `pkg`:

   - `models` contains all the data model declarations and a basic CRUD layer,
   - `rest`, `service` contains an implementation of the REST API,
   - `utils` has all the utility helpers,
 - `vendor` is the Go vendor directory,
 - `Gopkg.lock`, `Gopkg.toml` is the configuration of go-dep for package management,
 - the root directory contains mostly configuration files

## Code development guide
This project uses [`dep`](https://github.com/golang/dep) for dependency management,
 [`ginkgo`](https://github.com/onsi/ginkgo) for testing
and [`migrate`](https://github.com/golang-migrate/migrate) for running database
migrations. 

```bash
# Using go-dep to install dependency packages
dep ensure -v

# Using ginkgo for unit test
ginkgo -r -skipPackage=vendor --randomizeSuites --failOnPending --cover --trace --race --progress

# Or using the standard go test
go test -v ./...

# Using migrate for data migration
migrate -source file://./migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/tbox" up
```

In order to make setting it up easier, there's a `docker-compose.yml` , which contain the mysql image:

```bash
# Using docker-compose file to run mysql server localhost with port 3306
sudo docker-compose up -d
```

```bash
# Starting api 
go run main.go
```

```bash
# Send otp code to user phone number
curl -X POST \
  http://localhost:8080/users/otp/generate \
  -H 'Content-Type: application/json' \
  -d '{
	"phoneNumber":"0915558493"
}'
```

```bash
# Validate the otp with user phone number
curl -X POST \
  http://localhost:8080/users/otp/validate \
  -H 'Content-Type: application/json' \
  -d '{
	"phoneNumber":"0908280493",
	"otp": "123456"
}'
```