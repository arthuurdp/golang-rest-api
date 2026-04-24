package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"Hello_World/myapp/internal/adapter/handler"
	adapterrepo "Hello_World/myapp/internal/adapter/repositories"
	"Hello_World/myapp/internal/infrastructure/config"
	"Hello_World/myapp/internal/infrastructure/database"
	"Hello_World/myapp/internal/infrastructure/server"
	userusecase "Hello_World/myapp/internal/usecases/user"
)

func main() {
	cfg := config.Load()

	db, err := database.NewMySQLConnection(database.MySQLConfig{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	})
	if err != nil {
		log.Fatal("falha ao conectar no banco: ", err)
	}
	defer db.Close()

	userRepo := adapterrepo.NewMySQLUserRepository(db)

	createUserUC  := userusecase.NewCreateUserUseCase(userRepo)
	updateUsersUC := userusecase.NewUpdateUserUseCase(userRepo)
	getUserUC     := userusecase.NewGetUserUseCase(userRepo)
	getUsersUC    := userusecase.NewGetUsersUseCase(userRepo)

	userHandler := handler.NewUserHandler(createUserUC, updateUsersUC, getUserUC, getUsersUC)

	srv := server.NewServer(userHandler)
	srv.Setup()

	if err := srv.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("falha ao subir servidor: ", err)
	}
}