package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (m *MongoDb) GetAll() ([]types.Endpoint, error) {
	results, err := m.collection.Find(m.context, bson.D{})
	if err != nil {
		return nil, err
	}

	var endpoints []types.Endpoint
	for results.Next(m.context) {
		var document Document
		err = results.Decode(&document)
		if err != nil {
			continue
		}

		var endpoint types.Endpoint
		endpoint.Id = document.Id
		endpoint.HttpVerb = document.HttpVerb
		endpoint.Name = document.Name
		endpoint.Path = document.Path
		endpoint.Response = document.Response
		endpoints = append(endpoints, endpoint)
	}

	return endpoints, nil
}

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

func (m *MongoDb) Create(createEndpoint types.CreateEndpoint) error {
	_, err := m.collection.InsertOne(m.context, createEndpoint)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDb) Delete(id string) (types.Endpoint, error) {
	var endpoint types.Endpoint

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return endpoint, err
	}

	result := m.collection.FindOneAndDelete(m.context, bson.M{"_id": objectId})
	if result.Err() != nil {
		return endpoint, result.Err()
	}

	var document Document
	err = result.Decode(&document)
	if err != nil {
		return endpoint, err
	}

	endpoint.Id = document.Id
	endpoint.HttpVerb = document.HttpVerb
	endpoint.Name = document.Name
	endpoint.Path = document.Path

	return endpoint, nil
}
