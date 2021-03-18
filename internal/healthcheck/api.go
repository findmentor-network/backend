package healthcheck

import (
	. "github.com/findmentor-network/backend/pkg/echoextention/healthcheck"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func RegisterHandlers(instance *echo.Echo, mongoDb *mongo.Database) {

	opts := []Option{
		WithTimeout(2 * time.Second),
		WithChecker("mongoDb", CheckMongoDb(mongoDb)),
	}
	h := New(opts...).SetEndpoint("status")
	h.Use(instance)

}
