package apidockertest

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// When writing a fruit to the DB, we allow mongod to add the _id field
type FruitWritable struct {
	Name   string `bson:"name" json:"name" binding:"required"`
	Origin string `bson:"origin" json:"origin"`
	Price  int    `bson:"price" json:"price"`
}

// When reading, we recieve the _id field
type FruitReadable struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id"`
	Name   string             `bson:"name" json:"name"`
	Origin string             `bson:"origin" json:"origin"`
	Price  float64            `bson:"price" json:"price"`
}

//
func (f FruitReadable) Copy() FruitReadable {
	return FruitReadable{f.ID, f.Name, f.Origin, f.Price}
}
