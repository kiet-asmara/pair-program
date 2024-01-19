package model

import (
	"fmt"
	"time"
)

type Transaction struct {
	Id          string    `bson:"_id" json:"id"`
	Description string    `bson:"description" json:"description"`
	Amount      float64   `bson:"amount" json:"amount"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
}

func (t Transaction) Run() {
	fmt.Printf("Transaction run: %s, Amount: %.2f", t.Description, t.Amount)
}
