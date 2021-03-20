package mongohelper

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
)

func TestQueryBuilder(usecase *testing.T) {

	usecase.Run("When building query for mongo client", func(t *testing.T) {
		type args struct {
			k string
			v interface{}
		}
		_idTest, _ := primitive.ObjectIDFromHex("5d2399ef96fb765873a24bae")
		bsonTest := bson.M{}
		bsonTest["_id"] = _idTest
		tests := []struct {
			name      string
			args      args
			wantQuery bson.M
		}{
			{name: "With Empty Parameter should return nil", args: args{"", nil}, wantQuery: bson.M{}},
			{name: "With One Parameter should retun one key value pair", args: args{"Key", "Test"}, wantQuery: bson.M{"Key": "Test"}},
			{name: "With One Parameter should retun one key value pair", args: args{"Key", "Test"}, wantQuery: bson.M{"Key": "Test"}},
			{name: "With One Parameter and value is empty should return empty key", args: args{"Key", true}, wantQuery: bson.M{"Key": true}},
			//{name: "With id parameter should return key field is _id and value", args: map[string]string{"_id": "5d2399ef96fb765873a24bae"}, wantQuery: bsonTest},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				q := QueryOf()
				q.Add(tt.args.k, tt.args.v)
				if gotQuery := q.Build(); !reflect.DeepEqual(gotQuery, tt.wantQuery) {
					t.Errorf("BuildQuery() = %v, want %v", gotQuery, tt.wantQuery)
				}
			})
		}

	})
}
