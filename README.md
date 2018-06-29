# go-test
Golang Test Task

## Features
- Authorization and authentification via JWT
- Retrieving of users list depends on authorized user role
- Creation of a new user
- Deleting of a user

## Usage

> *Pay attention, docker and docker-compose must be installed to use commands below.*
>
> *All commands should be run in UNIX-like terminal.*

### Run Application
```shell
$ make run
```
HTTP sever will be started on port [:8080]

### Stop And Remove Application
```shell
$ make down
```

### Restart Application
```shell
$ make restart
```

## Dependencies
- [labstack/echo](https://github.com/labstack/echo)
- [dgrijalva/jwt-go](github.com/dgrijalva/jwt-go)
- [go-sql-driver/mysql](github.com/go-sql-driver/mysql)
- [jinzhu/configor](github.com/jinzhu/configor)
- [jmcvetta/randutil](github.com/jmcvetta/randutil)
- [mitchellh/mapstructure](github.com/mitchellh/mapstructure)
- [satori/go.uuid](github.com/satori/go.uuid)
- [thedevsaddam/govalidator](github.com/thedevsaddam/govalidator)
- [go.uber.org/dig](go.uber.org/dig)
