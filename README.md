[![Go](https://github.com/findmentor-network/backend/actions/workflows/go.yml/badge.svg)](https://github.com/findmentor-network/backend/actions/workflows/go.yml)

## Build Setup

```
go get github.com/findmentor-network/backend

make build

<$ ./backend
Findmentor API

Usage:
  backend [command]

Available Commands:
  api         find mentor api
  help        Help about any command

Flags:
  -c, --dbconn string   Database connection string (default "mongodb://root:example@127.0.0.1:27017/")
  -d, --dbname string   Database name (default "findmentor")
  -h, --help            help for backend
  -p, --port string     service port (default "5000")
  -t, --toggle          Help message for toggle

Use "backend [command] --help" for more information about a command.

```
or build & run on docker
```
make docker_build

docker run -it -p 5000:5000 backend api 
```

## Swagger

```
http://127.0.0.1:5000/swagger/index.html
```

## TODO
1. need to port generate.js to golang as command
