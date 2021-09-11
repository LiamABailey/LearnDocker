package internal

import (
  "github.com/LiamABailey/LearnDocker/internal/teststructs.go"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "context"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)


const (
  CONNSTR = "mongodb://localhost:27017"
  DB      = "food"
  FRUIT   = "fruit"
)

type MongoConnect struct {
  client    *mongo.Client
}


func newMongoClient() *MongoConnect {
  var err error
  cliops := M.Client().ApplyURI(CONNSTR)
  client, err := mongo.Connect(context.TODO(), cliops)
  M.client = client
  return err
}

func (M *MongoConnect) DisconnectMongoClient() error {
  err := M.client.Disconnect(context.TODO())
  if err != nil {
    return err
  }
  return nil
}

// adds a new Fruit to the db
func (M *MongoConnect) AddFruit(newfruit Fruit) error {
  coll = M.Database(DB).Collection(FRUIT)
  _, err = collection.InsertOne(context.TODO(), newfruit)
  if err != nil {
    return err
  }
  return nil
}

// adds a slice of fruits to the database
func (M *MongoConnect) AddFruitMulti(newfruits []Fruit) error {
    coll = M.Database(DB).Collection(FRUIT)
    _, err = collection.InsertMany(context.TODO(), newfruits)
    if err != nil {
      return err
    }
    return nil
}

// Returns fruit matching the provided ID.
func (M *MongoConnect) GetFruitByID(id primitive.ObjectID ) (FruitReadable, error) {
  var fruit FruitReadable
  filt := bson.D{primitive.E{Key: "_id", Value: id}}
  coll = M.Database(DB).Collection(FRUIT)
  err := collection.FindOne(context.TODO(), filt).Decode(&fruit)
  return FruitReadable, error
}

// Returns any fruit matching the provided name
func (M *MongoConnect) GetFruitByName(fruitname str) ([]FruitReadable, error) {
  var fruits []FruitReadable
  filt := bson.D{primitive.E{Key: "name", Value: fruitname}}
  coll = M.Database(DB).Collection(FRUIT)
  curs, err := episodesCollection.Find(context.TODO(), filt)
  if err != nil {
    return fruits, err
  }
  defer curs.Close(context.TODO())
  for curs.Next(context.TODO()){
    var fruit FruitReadable
    if err = curs.Decode(&fruit); err == nil {
      fruits = fruits.append(fruit.Copy())
    }
  }
  return fruits, err
}

// Remove fruit by ID
//func (M *MongoConnect) RemoveFruitByID(id primitive.ObjectID) (error){
//
//}
