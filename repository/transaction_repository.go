package repository

import (
	"context"
	"pair/model"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	ReadAll() ([]*model.Transaction, error)
	ReadID(transactionID int) (*model.Transaction, error)
	Update(transactionID string, input model.Transaction) error
	Delete(transactionID string) error
	DeleteAllBeforeMidnight() error
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

func (tr *transactionRepository) ReadAll() ([]*model.Transaction, error) {
	ctx := context.TODO()

	// get collection
	collection := tr.DB.Database("pair-program").Collection("transaction")

	var results []*model.Transaction

	//Passing the bson.D{{}} as the filter matches  documents in the collection
	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	//Finding multiple documents returns a cursor
	//Iterate through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {
		//Create a value into which the single document can be decoded
		var transaction *model.Transaction
		err := cur.Decode(&transaction)
		if err != nil {
			return nil, err
		}

		results = append(results, transaction)
	}

	return results, nil
}

func (tr *transactionRepository) ReadID(transactionID int) (*model.Transaction, error) {
	ctx := context.TODO()

	id := strconv.Itoa(transactionID)

	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	// get collection
	collection := tr.DB.Database("pair-program").Collection("transaction")

	var result model.Transaction

	//Passing the bson.D{{}} as the filter matches  documents in the collection
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (tr *transactionRepository) Update(transactionID string, input model.Transaction) error {
	ctx := context.TODO()

	filter := bson.D{primitive.E{Key: "_id", Value: transactionID}}

	update := bson.M{
		"$set": bson.M{
			"description": input.Description,
			"amount":      input.Amount,
		}}

	// get collection
	collection := tr.DB.Database("pair-program").Collection("transaction")

	//Passing the bson.D{{}} as the filter matches  documents in the collection
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (tr *transactionRepository) Delete(transactionID string) error {
	ctx := context.TODO()

	filter := bson.D{primitive.E{Key: "_id", Value: transactionID}}

	// get collection
	collection := tr.DB.Database("pair-program").Collection("transaction")

	//Passing the bson.D{{}} as the filter matches  documents in the collection
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (tr *transactionRepository) DeleteAllBeforeMidnight() error {
	ctx := context.TODO()

	// Mendapatkan waktu tengah malam hari ini
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())

	// Menetapkan kriteria pencarian untuk semua transaksi sebelum tengah malam
	filter := bson.D{{"createdAt", bson.D{{"$lt", midnight}}}}

	collection := tr.DB.Database("pair-program").Collection("transaction")

	// Menghapus semua dokumen dengan filter createdAt < midnight
	_, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
