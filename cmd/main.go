package main

import (
  "github.com/LiamABailey/LearnDocker/pkg/apidockertest"
)

func main(){
  mc, _ := apidockertest.NewMongoClient()
  srv := apidockertest.NewAPIServer(mc)
  srv.Run("localhost:8080")
  //fmt.Println(mc.GetFruitByName("orange"))
  //fw := apidockertest.FruitWritable{"pear","US",1}
  //posterr := mc.AddFruit(fw)
  //fmt.Println(posterr)
  //pears, _ := mc.GetFruitByName("pear")
  //fmt.Println(pears)
  //for _, p := range pears {
  //  fmt.Println(mc.DeleteFruitByID(p.ID))
  //}
  //fmt.Println(mc.DisconnectMongoClient())
}
