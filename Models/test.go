package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Test struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Test int                `bson:"test"`
}
