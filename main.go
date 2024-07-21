package main

import (
	_ "tokyokitten/docs"
	"tokyokitten/config"
	"tokyokitten/controller"
	"tokyokitten/helper"
	"tokyokitten/model"
	"tokyokitten/repository"
	"tokyokitten/router"
	"tokyokitten/service"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

// @title 	Kitten Service API
// @version	1.0
// @description A Kitten service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	// Database
	db := config.DBConnetion(&loadConfig)
	validate := validator.New()

	db.Table("kittens").AutoMigrate(&model.Kitten{})
	db.Table("users").AutoMigrate(&model.User{})

	// Init Repository
	userRepository:=repository.NewUsersRepositoryImpl(db)
	kittensRepository := repository.NewKittensRepositoryImpl(db)

	// Service
	authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)
	kittensService := service.NewKittensServiceImpl(kittensRepository, validate)

	// Controller
	kittensController := controller.NewKittensController(kittensService)
	authenticationController := controller.NewAuthenticationController(authenticationService)
	usersController := controller.NewUsersController(userRepository)

	// Router
	routes := router.NewRouter(userRepository, 
		authenticationController, 
		usersController, 
		kittensController)

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err:= server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
