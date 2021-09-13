package datastore

import (
	"context"
	"fmt"
	"time"

	backoff "github.com/cenkalti/backoff/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collections = struct {
	Servers string
}{
	"servers",
}

func Connect(URI string) (*mongo.Client, error) {
	var client *mongo.Client
	connect := func() error {
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		client, err = mongo.Connect(ctx, options.Client().ApplyURI(URI))
		if err != nil {
			fmt.Printf("unable to connect to datastore: %s\n", err.Error())
		}

		return err
	}

	if err := backoff.Retry(connect, backoff.NewExponentialBackOff()); err != nil {
		return nil, err
	}

	return client, nil
}
