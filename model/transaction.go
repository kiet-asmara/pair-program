package model

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Id          primitive.ObjectID `bson:"_id"`
	Description string             `bson:"description"`
	Amount      float64            `bson:"amount"`
}

func (t Transaction) Run() {
	fmt.Printf("Transaction run: %s, Amount: %.2f", t.Description, t.Amount)
}
