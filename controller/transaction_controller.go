package controller

import (
	"pair/model"
	"pair/repository"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransactionController(transactionRepository repository.TransactionRepository) *TransactionController {
	return &TransactionController{
		TransactionRepository: transactionRepository,
	}
}

func (t *TransactionController) CreateTranscation(c echo.Context) error {
	var newTranscation model.Transaction
	if err := c.Bind(&newTranscation); err != nil {
		return c.JSON(400, echo.Map{
			"message": "invalid request",
		})
	}

	if err := t.TransactionRepository.Create(&newTranscation); err != nil {
		return c.JSON(500, echo.Map{
			"message": "failed to create transaction",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "transaction created successfully",
	})
}
