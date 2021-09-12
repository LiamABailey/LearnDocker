package apidockertest

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type APIServer struct {
  connector *MongoConnect
  router    *gin.Engine
}

func NewAPIServer(mc *MongoConnect) *APIServer {
  svr := &APIServer{connector: mc}
  svr.router = gin.Default()

  svr.router.GET("/fruits/name/:name", svr.getFruitByName)
  svr.router.GET("/fruits/id/:id", svr.getFruitByID)
  svr.router.POST("/fruits", svr.addFruit)
  svr.router.DELETE("/fruits/id/:id", svr.deleteFruitByID)

  return svr
}

func (srv *APIServer) Run(address string) error {
  return srv.router.Run(address)
}

func (srv *APIServer) getFruitByName(ctx *gin.Context) {
  fname := ctx.Param("name")
  fruits, err := srv.connector.GetFruitByName(fname)
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, errorResponse(err))
    return
  }
  ctx.IndentedJSON(http.StatusOK, fruits)
}

func (srv *APIServer) getFruitByID(ctx *gin.Context) {
  id, _ := idFromContext(ctx)
  fruit, err := srv.connector.GetFruitByID(id)
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, errorResponse(err))
    return
  }
  ctx.IndentedJSON(http.StatusOK, fruit)
}

func (srv *APIServer) addFruit(ctx *gin.Context) {
  var fruit FruitWritable
  if err := ctx.ShouldBindJSON(&fruit); err != nil {
    ctx.JSON(http.StatusBadRequest, errorResponse(err))
    return
  }

  err := srv.connector.AddFruit(fruit)
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, errorResponse(err))
    return
  }
  ctx.JSON(http.StatusCreated, nil)
}

func (srv *APIServer) deleteFruitByID(ctx *gin.Context) {
    id, _ := idFromContext(ctx)
    dr, err := srv.connector.DeleteFruitByID(id)
    if err != nil {
      ctx.JSON(http.StatusInternalServerError, errorResponse(err))
      return
    }
    ctx.IndentedJSON(http.StatusOK, dr)
}

func errorResponse(err error) gin.H {
  return gin.H{"error": err.Error()}
}

func idFromContext(ctx * gin.Context) (primitive.ObjectID,error){
  return primitive.ObjectIDFromHex(ctx.Param("id"))
}
