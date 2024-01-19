package controller

import (
	"pair/model"

	"github.com/labstack/echo"
)

type TransactionController struct {
}

func (t *TransactionController) CreateTranscation(c echo.Context) error {
	var newTranscation model.Transaction
	if err := c.Bind(&newTranscation); err != nil {
		return c.JSON(400, echo.Map{
			"message": "invalid request",
		})
	}

	if err := t.TrascationRepository.Create(&newTranscation); err != nil {
		return c.JSON(500, echo.Map{
			"message": "failed to create transaction",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "transaction created successfully",
	})
}
