package person

import (
	"context"
	"github.com/findmentor-network/backend/pkg/errors"
	"github.com/findmentor-network/backend/pkg/log"
	mongohelper "github.com/findmentor-network/backend/pkg/mongoextentions"
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
func (m mongoRepository) Get(ctx context.Context, query *mongohelper.Query, pg *pagination.Pages) (result []*Person, err error) {

	batchSize := int32(pg.PerPage)
	skip := int64(pg.Offset())
	limit := int64(pg.Limit())
	opt := &options.FindOptions{
		BatchSize: &batchSize,
		Limit:     &limit,
		Skip:      &skip,
	}
	if len(pg.Sort) > 0 {
		intSortBy := 1
		if pg.SortBy == "desc" {
			intSortBy = -1
		}
		opt.SetSort(D{{pg.Sort, intSortBy}})
	}

	c, err := m.db.Find(ctx, query.Build(), opt)

	if err != nil {
		return nil, errors.New(DatabaseError, err.Error())
	}
	for c.Next(ctx) {
		var elem *Person
		err = c.Decode(&elem)
		if err != nil {
			log.Logger.Errorf("failed to get cursor. %s", err.Error())
		} else {
			result = append(result, elem)
		}
	}

	return
}
