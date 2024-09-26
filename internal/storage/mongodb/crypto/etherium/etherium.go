package rmongo

import (
	"context"
	"gihub.com/robbitancor/simple-microservice/internal/domain/crypto/etherium"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoRepository struct {
	client *mongo.Client
}

func NewEtheriumMongoRepository(client *mongo.Client) *MongoRepository {
	return &MongoRepository{client}
}

func (e *MongoRepository) Create(eth etherium.Etherium) error {
	db := e.client.Database("testing")
	col := db.Collection("crypto")
	_, err := col.InsertOne(context.TODO(), bson.D{
		{
			Key: "amount", Value: eth.Balance,
		},
	})

	if err != nil {
		log.Error(err)
	}
	return nil
}

func (e *MongoRepository) Read(eth etherium.Etherium) (etherium.Etherium, error) {
	return etherium.Etherium{}, nil
}

func (e *MongoRepository) Update(eth etherium.Etherium) error {
	return nil
}

func (e *MongoRepository) Delete(eth etherium.Etherium) error {
	return nil
}
