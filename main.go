package main

import (
	"fmt"
	"log"
	"pair/config"
	"pair/controller"
	"pair/repository"
	"pair/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := config.InitDB()
	if err != nil {
		fmt.Println("failed connecting to db")
		log.Fatal(err)
	}

	TransactionRepository := repository.NewTransactionRepository(db)

	TransactionController := controller.NewTransactionController(TransactionRepository)

	router.RegisterRoutes(e, TransactionController)

	e.Start(":8081")
}
