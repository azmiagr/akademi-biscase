package main

import (
	"akademi-business-case/internal/handler/rest"
	"akademi-business-case/internal/repository"
	"akademi-business-case/internal/service"
	"akademi-business-case/pkg/bcrypt"
	"akademi-business-case/pkg/config"
	"akademi-business-case/pkg/database/mariadb"
	"akademi-business-case/pkg/jwt"
	"log"
)

func main() {
	config.LoadEnvironment()

	db, err := mariadb.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	err = mariadb.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	bcrypt := bcrypt.Init()
	jwt := jwt.Init()
	svc := service.NewService(repo, bcrypt, jwt)

	r := rest.NewRest(svc)
	r.MountEndpoint()
	r.Run()
}
