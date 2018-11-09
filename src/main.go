package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var (
	mongoBookingsCollection *mgo.Collection
	mongoLogsCollection     *mgo.Collection
)

type jsonType gin.H

func init() {
	mongoSession, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		log.Fatal("Unable to connect to the data store: ", err.Error())
	}
	mongoBookingsCollection = mongoSession.DB(os.Getenv("MONGO_DB_NAME")).C(os.Getenv("MONGO_DB_BOOKINGS_COLLECTION"))
	mongoLogsCollection = mongoSession.DB(os.Getenv("MONGO_DB_NAME")).C(os.Getenv("MONGO_DB_LOGS_COLLECTION"))
}

func main() {
	router := gin.Default()
	router.GET("/booking/:objectId", handleGet)
	router.POST("/booking", handleUpsert)
	router.POST("/booking/:objectId", handleUpsert)
	router.Run(os.Getenv("APP_PORT"))
}

func handleGet(c *gin.Context) {
	var (
		objectId bson.ObjectId = bson.ObjectIdHex(c.Param("objectId"))
	)
	result, err := mongoGetByObjectId(objectId)
	if err != nil {
		handleError(c, err)
		return
	}
	handleSuccess(c, jsonType{"booking": result})
}

func handleUpsert(c *gin.Context) {
	var (
		objectId string = c.Param("objectId")
		body     jsonType
		id       bson.ObjectId
	)
	c.BindJSON(&body)
	body["updatedAt"] = time.Now()
	if objectId == "" {
		id = bson.NewObjectId()
		body["createdAt"] = time.Now()
	} else {
		id = bson.ObjectIdHex(objectId)
	}
	change := bson.M{"$set": body}
	_, err := mongoBookingsCollection.UpsertId(id, change)
	if err != nil {
		handleError(c, err)
		return
	}
	_ = mongoLogsCollection.Insert(jsonType{"booking_id": id.Hex(), "body": body})
	updatedDocument, _ := mongoGetByObjectId(id)
	handleSuccess(c, jsonType{"data": updatedDocument})
}

func mongoGetByObjectId(objectId bson.ObjectId) (interface{}, error) {
	var result interface{}
	if err := mongoBookingsCollection.FindId(objectId).One(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func handleError(c *gin.Context, e error) {
	// fmt.Printf("%+v\n", e.Error())
	c.JSON(http.StatusUnauthorized, jsonType{"error": e.Error()})
}

func handleSuccess(c *gin.Context, data jsonType) {
	// fmt.Printf("%+v\n", data)
	c.JSON(http.StatusOK, jsonType(data))
}
