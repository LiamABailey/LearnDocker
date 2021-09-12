package main

import (
  "fmt"
  "github.com/LiamABailey/LearnDocker/pkg/apidockertest"
)

func main(){
  mc, _ := apidockertest.NewMongoClient()
  fmt.Println(mc.GetFruitByName("orange"))
  fmt.Println(mc.DisconnectMongoClient())
}
