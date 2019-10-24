package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	mongoBookingsCollection *mongo.Collection
	mongoLogsCollection     *mongo.Collection
)

type jsonType gin.H

func init() {
	db, err := GetDB()
	if err != nil {
		log.Fatal("Unable to connect to the data store: ", err.Error())
	}
	mongoBookingsCollection = db.Collection(os.Getenv("MONGO_DB_BOOKINGS_COLLECTION"))
	mongoLogsCollection = db.Collection(os.Getenv("MONGO_DB_LOGS_COLLECTION"))
}

func main() {
	router := gin.Default()
	router.GET("/status", handleGetStatus)
	router.GET("/booking/:objectId", handleGet)
	// router.POST("/booking", handleUpsert)
	// router.POST("/booking/:objectId", handleUpsert)
	router.Run(os.Getenv("APP_PORT"))
}

func handleGetStatus(c *gin.Context) {
	handleSuccess(c, jsonType{"status": "OK"})
}

func handleGet(c *gin.Context) {
	result, err := GetByObjectID(c.Param("objectId"), mongoBookingsCollection)
	if err != nil {
		handleError(c, err)
		return
	}
	handleSuccess(c, jsonType{"booking": result})
}

// func handleUpsert(c *gin.Context) {
// 	var (
// 		objectId = c.Param("objectId")
// 		body     jsonType
// 		id       bson.ObjectId
// 	)
// 	c.BindJSON(&body)
// 	body["updatedAt"] = time.Now()
// 	if objectId == "" {
// 		id = bson.NewObjectId()
// 		body["createdAt"] = time.Now()
// 	} else {
// 		id = bson.ObjectIdHex(objectId)
// 	}
// 	change := bson.M{"$set": body}
// 	_, err := mongoBookingsCollection.UpsertId(id, change)
// 	if err != nil {
// 		handleError(c, err)
// 		return
// 	}
// 	_ = mongoLogsCollection.Insert(jsonType{"booking_id": id.Hex(), "body": body})
// 	updatedDocument, _ := GetByObjectID(id.String(), mongoBookingsCollection)
// 	handleSuccess(c, jsonType{"data": updatedDocument})
// }

func handleError(c *gin.Context, e error) {
	// fmt.Printf("%+v\n", e.Error())
	c.JSON(http.StatusUnauthorized, jsonType{"error": e.Error()})
}

func handleSuccess(c *gin.Context, data jsonType) {
	// fmt.Printf("%+v\n", data)
	c.JSON(http.StatusOK, jsonType(data))
}
