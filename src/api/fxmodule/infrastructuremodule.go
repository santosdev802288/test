// Copyright 2020 Siigo. All rights reserved.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package fxmodule

import (
	"context"
	"siigo.com/kubgo/src/api/config"
	"time"

	"github.com/sony/gobreaker"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"siigo.com/kubgo/src/infrastructure/finder"
	"siigo.com/kubgo/src/infrastructure/repository"
)

// InfrastructureModule Infrastructure Module Finders and Repositories
var InfrastructureModule = fx.Options(
	fx.Provide(
		NewMongoClient,
		NewMongoContext,
		NewKubgoCollection,

		repository.NewKubgoRepository,
		finder.NewKubgoFinder,
	),
)

// NewMongoClient Create Mongo Connection
func NewMongoClient(config *config.Configuration) *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(config.Mongo.ConnectionString))

	if err != nil {
		panic(err.Error())
	}
	return client
}

// NewMongoContext Create New Mongo Context
func NewMongoContext() context.Context {
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	return ctx
}

// NewKubgoCollection Connect to Mongo Collection
func NewKubgoCollection(ctx context.Context, client *mongo.Client, config *config.Configuration) *mongo.Collection {
	err := client.Connect(ctx)
	if err != nil {
		panic(err.Error())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil
	}

	collection := client.
		Database(config.Mongo.Database).
		Collection(config.Mongo.Collection)

	// create mongo indexes
	if _, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "_id", Value: 1}},
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{Keys: bson.D{{Key: "occurred_at", Value: -1}}},
	}); err != nil {
		panic(err)
	}

	return collection
}

func NewCircuitBreaker() *gobreaker.CircuitBreaker {
	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name: "circuit-breaker",
		//Timeout: gobreaker.DefaultTimeout,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		},
	})
}
