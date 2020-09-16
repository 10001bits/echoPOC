package types

import "go.mongodb.org/mongo-driver/mongo"

type Coll struct{
	UserCollection *mongo.Collection
}
