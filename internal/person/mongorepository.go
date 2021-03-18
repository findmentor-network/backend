package person

import (
	"context"
	"github.com/findmentor-network/backend/pkg/errors"
	"github.com/findmentor-network/backend/pkg/log"
	"github.com/findmentor-network/backend/pkg/pagination"
	. "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	db *mongo.Collection
}

const collectionName = "persons"

func NewRepository(database *mongo.Database) Repository {

	return &mongoRepository{
		db: database.Collection(collectionName),
	}
}
func (m mongoRepository) Get(ctx context.Context, pg *pagination.Pages) (result []*Person, err error) {

	batchSize := int32(pg.PerPage)
	skip := int64(pg.Offset())
	limit := int64(pg.Limit())
	opt := &options.FindOptions{
		BatchSize: &batchSize,
		Limit:     &limit,
		Skip:      &skip,
	}
	c, err := m.db.Find(ctx, M{}, opt)
	//todo:better error handling
	if err != nil {
		return nil, errors.New(DatabaseError, err.Error())
	}
	for c.Next(ctx) {
		var elem *Person
		err = c.Decode(&elem)
		if err != nil {
			log.Logger.Error("failed to get cursor. %s", err.Error())
		} else {
			result = append(result, elem)
		}
	}

	return
}
