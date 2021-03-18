package mongohelper

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
)

func TestNewDatabase(usecase *testing.T) {
	usecase.Run("When creating a new mongo client", func(t *testing.T) {

		tests := []struct {
			name     string
			args     []string
			wantErr  bool
			whichErr error
		}{
			{name: "with valid parameters_should not return error", args: []string{"mongodb://root:example@localhost:27017", "admin"}, wantErr: false},
			{name: "with invalid parameters_should return error", args: []string{"mongodb://an_invalid_address:27017?connectTimeoutMS=300", "admin"}, wantErr: true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := NewDatabase(tt.args[0], tt.args[1])

				if err != nil && tt.wantErr == true && errors.Cause(err) != ConnectionError {
					t.Errorf("NewDatabase() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	})

	usecase.Run("When building query for mongo client", func(t *testing.T) {

		_idTest, _ := primitive.ObjectIDFromHex("5d2399ef96fb765873a24bae")
		bsonTest := bson.M{}
		bsonTest["_id"] = _idTest
		tests := []struct {
			name      string
			args      map[string]string
			wantQuery bson.M
		}{
			{name: "With Empty Parameter should return nil", args: nil, wantQuery: bson.M{}},
			{name: "With One Parameter should retun one key value pair", args: map[string]string{"Key": "Test"}, wantQuery: bson.M{"Key": "Test"}},
			{name: "With One Parameter and value is empty should return empty key", args: map[string]string{"Key": ""}, wantQuery: bson.M{}},
			//{name: "With id parameter should return key field is _id and value", args: map[string]string{"_id": "5d2399ef96fb765873a24bae"}, wantQuery: bsonTest},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if gotQuery := BuildQuery(tt.args); !reflect.DeepEqual(gotQuery, tt.wantQuery) {
					t.Errorf("BuildQuery() = %v, want %v", gotQuery, tt.wantQuery)
				}
			})
		}

	})
}
