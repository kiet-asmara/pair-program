package repository

import (
	"context"
	"pair/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	// ReadAll(transactionID int) ([]*model.Transaction, error)
	// ReadID(transactionID int) (*model.Transaction, error)
	// Update(transactionID int) (*model.Transaction, error)
	// Delete(transactionID int) (*model.Transaction, error)
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
	collection := tr.DB.Database("pair-program").Collection("transaction")

	// insert db
	_, err := collection.InsertOne(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}
