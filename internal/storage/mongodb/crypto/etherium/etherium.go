package rmongo

import (
	"gihub.com/robbitancor/simple-microservice/internal/domain/crypto/etherium"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoRepository struct {
	client *mongo.Client
}

func NewEtheriumMongoRepository(client *mongo.Client) *MongoRepository {
	return &MongoRepository{client}
}

func (e *MongoRepository) Create(eth etherium.Etherium) error {
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
