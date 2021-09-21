package apidockertest

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type APIServer struct {
	connector *MongoConnect
	router    *gin.Engine
}

func NewAPIServer(mc *MongoConnect) *APIServer {
	svr := &APIServer{connector: mc}
	svr.router = gin.Default()
	//route GET, POST, and DELETE methods
	svr.router.GET("/fruits/name/:name", svr.getFruitByName)
	svr.router.GET("/fruits/id/:id", svr.getFruitByID)
	svr.router.POST("/fruits", svr.addFruit)
	svr.router.DELETE("/fruits/id/:id", svr.deleteFruitByID)

	return svr
}

func (srv *APIServer) Run(address string) error {
  return srv.router.Run(address)
}

// Search a Fruit by its name
// Gets 0 or more fruit with matching name
func (srv *APIServer) getFruitByName(ctx *gin.Context) {
	fname := ctx.Param("name")
	fruits, err := srv.connector.GetFruitByName(fname)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	//send a 200 OK status to suggest successful GET, plus
	//the data requested
	ctx.IndentedJSON(http.StatusOK, fruits)
}

// Search a fruit by its ID
// Collects 0 or 1 fruit with matching ID
func (srv *APIServer) getFruitByID(ctx *gin.Context) {
	id, _ := idFromContext(ctx)
	fruit, err := srv.connector.GetFruitByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	//send a 200 OK status to suggest successful GET, plus
	//the data requested
	ctx.IndentedJSON(http.StatusOK, fruit)
}

// Add a fruit
func (srv *APIServer) addFruit(ctx *gin.Context) {
	var fruit FruitWritable
	// bind the JSON body into a FruitWritable struct, leveraging json tags
	if err := ctx.ShouldBindJSON(&fruit); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := srv.connector.AddFruit(fruit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	//send a 201 created status back to the user to show successful POST
	ctx.JSON(http.StatusCreated, nil)
}

// Delete up to one fruit by ID
func (srv *APIServer) deleteFruitByID(ctx *gin.Context) {
	id, _ := idFromContext(ctx)
	dr, err := srv.connector.DeleteFruitByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	//send a 200 OK status to suggest successful deletion
	ctx.IndentedJSON(http.StatusOK, dr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

//create the mongo ObjectID primitive from the hex-id string
func idFromContext(ctx *gin.Context) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(ctx.Param("id"))
}
