package mongohelper

import (
	"context"
	"github.com/findmentor-network/backend/pkg/log"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewDatabase(uri, databaseName string) (db *mongo.Database, err error) {

	log.Logger.Infof("Mongo:Connection Uri:%s", uri)
	clientOptions := options.
		Client().
		ApplyURI(uri)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {

		return db, errors.Wrap(ConnectionError, err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Logger.Errorf("Mongo: mongo client couldn't connect with background context: %v", err)
		return db, errors.Wrap(ConnectionError, err.Error())
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return db, errors.Wrap(PingError, err.Error())
	}

	db = client.Database(databaseName)

	return db, err
}
