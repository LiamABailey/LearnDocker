package internal

import (
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type Fruit struct {
  ID      primitive.ObjectID  `bson:"_id"`
  Name    string              `bson:"name"`
  Origin  string              `bson:"origin"`
  Price   int                 `bson:"number"`
}
