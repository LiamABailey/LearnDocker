package apidockertest

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

// define system environment variable names
const (
	CONNSTR    = "MONGOCONN"
	DB         = "MONGODB"
	COLLECTION = "MONGOCOLLECTION"
)

// Client-containing struct, supports method attachment for cleaner code.
type MongoConnect struct {
	Client *mongo.Client
}

// Create a new Mongo client
func NewMongoConnect() (*MongoConnect, error) {
	M := MongoConnect{}
	cliops := options.Client().ApplyURI(os.Getenv(CONNSTR))
	client, err := mongo.Connect(context.TODO(), cliops)
	M.Client = client
	return &M, err
}

// Disconnect the mongo client
func (M *MongoConnect) DisconnectMongoClient() error {
	err := M.Client.Disconnect(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

// Adds a new Fruit to the db
func (M *MongoConnect) AddFruit(newfruit FruitWritable) error {
	coll := M.getCollection()
	_, err := coll.InsertOne(context.TODO(), newfruit)
	if err != nil {
		return err
	}
	return nil
}

// Returns fruit matching the provided ID.
func (M *MongoConnect) GetFruitByID(id primitive.ObjectID) (FruitReadable, error) {
	var fruit FruitReadable
	// filter by the passed ID
	filt := bson.M{"_id": id}
	coll := M.getCollection()
	err := coll.FindOne(context.TODO(), filt).Decode(&fruit)
	return fruit, err
}

// Returns any fruit matching the provided name
func (M *MongoConnect) GetFruitByName(fruitname string) ([]FruitReadable, error) {
	var fruits []FruitReadable
	// filter by the provided fruit name
	filt := bson.M{"name": fruitname}
	coll := M.getCollection()
	curs, err := coll.Find(context.TODO(), filt)
	if err != nil {
		return fruits, err
	}
	defer curs.Close(context.TODO())
	// Decode each fruit returned
	for curs.Next(context.TODO()) {
		var fruit FruitReadable
		if err = curs.Decode(&fruit); err == nil {
			fruits = append(fruits, fruit.Copy())
		}
	}
	return fruits, err
}

// Deletes a fruit matching the provided ID
func (M *MongoConnect) DeleteFruitByID(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	// filter by the provided ID
	filt := bson.M{"_id": id}
	coll := M.getCollection()
	dr, err := coll.DeleteOne(context.TODO(), filt)
	return dr, err
}

// connector to collections
func (M *MongoConnect) getCollection() *mongo.Collection {
	return M.Client.Database(os.Getenv(DB)).Collection(os.Getenv(COLLECTION))
}
