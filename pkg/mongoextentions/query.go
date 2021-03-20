package mongohelper

import (
	. "go.mongodb.org/mongo-driver/bson"
	"strings"
)

type Query struct {
	params map[string]interface{}
}

func QueryOf() *Query {
	return &Query{params: map[string]interface{}{}}
}
func (receiver *Query) Add(k string, v interface{}) {

	if v == nil {
		return
	}
	if _, ok := receiver.params[k]; !ok {
		receiver.params[k] = v
	}
}
func (receiver *Query) Build() (query M) {

	query = M{}
	for field, value := range receiver.params {

		if value == nil {
			continue
		}
		switch c := value.(type) {
		case string:
			if len(c) == 0 {
				continue
			}
			if strings.Contains(c, ",") {
				query[field] = M{"$in": strings.Split(c, ",")}
			} else {
				query[field] = c
			}
			break
		default:
			query[field] = c
		}
	}
	return query
}
