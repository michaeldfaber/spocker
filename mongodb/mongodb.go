package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	types "spocker/types"
)

type MongoDb struct {
	collection *mongo.Collection
	context    context.Context
	client     *mongo.Client
	cancel     context.CancelFunc
}

func New() (*MongoDb, error) {
	var mongodb MongoDb

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		return &mongodb, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		cancel()
		return &mongodb, err
	}

	collection := client.Database("spocker").Collection("endpoints")

	mongodb.collection = collection
	mongodb.context = ctx
	mongodb.client = client
	mongodb.cancel = cancel

	return &mongodb, nil
}

func (m *MongoDb) Disconnect() {
	m.client.Disconnect(m.context)
	m.cancel()
}

func (m *MongoDb) GetAll() {}

func (m *MongoDb) GetResponse(httpVerb string, name string) (interface{}, error) {
	var filter Filter
	filter.httpVerb = httpVerb
	filter.name = name

	result := m.collection.FindOne(m.context, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var document Document
	result.Decode(&document)
	return document.Response, nil
}

func (m *MongoDb) Create(endpoint types.Endpoint) error {
	_, err := m.collection.InsertOne(m.context, endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDb) Delete(deleteEndpointRequest types.DeleteEndpointRequest) {}
