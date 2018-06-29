package main

import (
	"go-test/src/domain"
	"go-test/src/infrastructure"
	"go-test/src/interfaces/repositories"
	"go-test/src/interfaces/webservices"
	"go-test/src/usecases"

	"go.uber.org/dig"
)

func buildDIContainer() *dig.Container {
	container := dig.New()

	checkErr(container.Provide(loadConfig))
	checkErr(container.Provide(newDbWrapper))
	checkErr(container.Provide(newLogger))
	checkErr(container.Provide(newJWTHandler))
	checkErr(container.Provide(newUserRepository))
	checkErr(container.Provide(usecases.NewAuthInteractor))
	checkErr(container.Provide(usecases.NewUserInteractor))
	checkErr(container.Provide(webservices.NewAuthWebservice))
	checkErr(container.Provide(webservices.NewUsersWebservice))
	checkErr(container.Provide(NewServer))

	return container
}

func newDbWrapper(c *config) repositories.DbHandler {
	return infrastructure.NewMySQLHandler(c.DBSource)
}

func newUserRepository(db repositories.DbHandler) domain.UserRepository {
	return repositories.NewUserRepository(db)
}

func newLogger(c *config) webservices.Logger {
	return infrastructure.NewLogger(c.Debug)
}

func newJWTHandler(c *config) webservices.JWTHandler {
	return infrastructure.NewJWTHandler(c.SigningKey, c.LifeTime)
}
