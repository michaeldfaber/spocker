package mongodb

import (
	"context"
	"fmt"
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
		fmt.Printf("1 %v\n", err)
		return &mongodb, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		cancel()
		fmt.Printf("2 %v\n", err)
		return &mongodb, err
	}
	// defer client.Disconnect(ctx)

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

func (m *MongoDb) Create(createEndpointRequest types.CreateEndpointRequest) error {
	var endpoint types.Endpoint
	endpoint.HttpVerb = createEndpointRequest.HttpVerb
	endpoint.Name = createEndpointRequest.Endpoint
	endpoint.Path = createEndpointRequest.Endpoint
	endpoint.Response = createEndpointRequest.ExpectedJsonResponse

	_, err := m.collection.InsertOne(m.context, endpoint)
	if err != nil {
		fmt.Printf("3 %v\n", err)
		return err
	}

	return nil
}

func (m *MongoDb) Delete(deleteEndpointRequest types.DeleteEndpointRequest) {}
