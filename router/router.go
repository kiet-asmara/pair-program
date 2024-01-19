package router

import (
	"pair/controller"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, transactionController *controller.TransactionController) {
	e.POST("/transactions", transactionController.CreateTranscation)
	e.GET("/transactions", transactionController.GetAllTransaction)
	e.GET("/transactions/:id", transactionController.GetTransactionByID)
	e.PUT("/transactions/:id", transactionController.UpdateTransaction)
	e.DELETE("/transactions/:id", transactionController.Deletetransaction)
}
