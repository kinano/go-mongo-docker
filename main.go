package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// @todo @kinano Finalize the data structs for a booking
type Person struct {
	firstName string
	lastName  string
	email     string
	phone     string
}

type Stop struct {
	name             string
	leavesAt         time.Time
	arrivesAt        time.Time
	latitude         float64
	longitude        float64
	formattedAddress string
	group            string
	timeZone         string
}

type TravelEstimate struct {
	pickup   Stop
	dropoff  Stop
	distance int
	duration int
}

type Booking struct {
	ID              bson.ObjectId `bson:"_id,omitempty"`
	name            string
	owner           Person
	riders          []Person
	stops           []Stop
	travelEstimates []TravelEstimate
}

type jsonType gin.H

func init() {
	mongoSession, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}
	mongoBookingsCollection = mongoSession.DB(os.Getenv("MONGO_DB_NAME")).C(os.Getenv("MONGO_DB_BOOKINGS_COLLECTION"))
	mongoLogsCollection = mongoSession.DB(os.Getenv("MONGO_DB_NAME")).C(os.Getenv("MONGO_DB_LOGS_COLLECTION"))
}

func main() {
	router := gin.Default()
	router.GET("/api/booking/:objectId", handleGet)
	router.POST("/api/booking", handleUpsert)
	router.POST("/api/booking/:objectId", handleUpsert)
	router.POST("/api/login", handleLogin)
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

func handleLogin(c *gin.Context) {
	var (
		body jsonType
	)
	c.BindJSON(&body)
	// Ruby expects {"user": {}}
	body = jsonType{
		"user": body,
	}

	bytesRepresentation, err := json.Marshal(body)
	log.Printf("bytes %+v\n", bytesRepresentation)

	if err != nil {
		handleError(c, err)
	}

	// @todo @kinano Move url to config
	resp, err := http.Post("http://localhost:3000/users/sign_in", "application/json", bytes.NewBuffer(bytesRepresentation))
	// log.Printf("Login response cookies: %+v\n", resp.Cookies())
	if err != nil {
		handleError(c, err)
	}
	for _, cookie := range resp.Cookies() {
		fmt.Println("Found a cookie named:", cookie.Name)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	handleSuccess(c, jsonType{"data": result})

}
