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
