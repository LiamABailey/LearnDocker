package internal

const (
  CONNSTR = "mongodb://localhost:27017"
  DB      = "food"
  FRUIT   = "fruit"
)

type MongoConnect struct {

}


func (M *MongoConnect) GetMongoClient() error {

}

func (M *MongoConnect) DisconnectMongoClient() error {

}
