package initialize

import "go.mongodb.org/mongo-driver/bson"

var Mongo = new(mongo)

type (
	mongo struct{}
	Index struct {
		V    any      `bson:"v"`
		Ns   any      `bson:"ns`
		Key  []bson.E `bson:"key`
		Name string   `bson:"name"`
	}
)

func (m *mongo) Initialization() error {
	return nil
}
