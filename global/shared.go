package global

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DBConnection   *mongo.Client
	BlogCollection *mongo.Collection
)
