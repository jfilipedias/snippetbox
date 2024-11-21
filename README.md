# Snippetbox

A Web Application build with Go.

## Getting started

### Setup the HTTPS connection

To run the server with a local HTTPS connection a TLS certificate is needed. To get a certificate for development environment you can use the `crypt/tls` Go package. It has the `generate_cert.go` tool to generate a self-signed certificate. To use this tool you can run the following commands:

```sh
$ mkdir tls
$ cd tls
$ go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

### Setup the database container

The web app uses a MySQL database to persist the users data. You can setup a docker container running a MySQL image with a clean database with the following command:

```sh
$ docker compose up -d
```

## Dependencies

- [alexedwards/scs](https://github.com/alexedwards/scs): A session manager tool to handle flash messages.
- [go-playground/form](https://github.com/go-playground/form): A form parser to get the fields value.
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql): A MySQL driver.
- [justinas/alice](https://github.com/justinas/alice): A middleware chaining tool.
