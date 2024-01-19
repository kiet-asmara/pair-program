package repository

import (
	"context"
	"fmt"
	"pair/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	ReadAll(transactionID int) ([]*model.Transaction, error)
	ReadID(transactionID int) (*model.Transaction, error)
	Update(transactionID int) (*model.Transaction, error)
	Delete(transactionID int) (*model.Transaction, error)
}

type transactionRepository struct {
	DB *mongo.Client
}

func NewTransactionRepository(db *mongo.Client) *transactionRepository {
	return &transactionRepository{
		DB: db,
	}
}

func (tr *transactionRepository) Create(transaction *model.Transaction) error {
	ctx := context.TODO()

	// get collection
	collection := mongoClient.Database("").Collection("")

	// insert db
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		fmt.Println(err)
	}

}
