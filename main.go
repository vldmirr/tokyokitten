package main

import (
	"tokyokitten/config"
	"tokyokitten/controller"
	"tokyokitten/helper"
	"tokyokitten/model"
	"tokyokitten/repository"
	"tokyokitten/router"
	"tokyokitten/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DBConnetion()
	validate := validator.New()

	db.Table("kittens").AutoMigrate(&model.Kitten{})

	// Repository
	kittensRepository := repository.NewkittensRepositoryImpl(db)

	// Service
	kittensService := service.NewKittensServiceImpl(kittensRepository, validate)

	// Controller
	kittensController := controller.NewkittensController(kittensService)

	// Router
	routes := router.NewRouter(kittensController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
