package main

import (
	"go-test/src/domain"
	"go-test/src/interfaces/repositories"
)

func main() {
	container := buildDIContainer()
	err := container.Invoke(func(server *Server, db repositories.DbHandler, ur domain.UserRepository) {
		// Close db connection
		defer db.Close()

		// Run migrations and seed the bootstrap data to the database
		checkErr(repositories.Migrate(db))
		checkErr(repositories.Seed(ur))

		// Run server
		server.Run()
	})
	if err != nil {
		panic(err)
	}
}
