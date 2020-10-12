package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(s string) *mongo.Collection {

	dbUser := os.Getenv("USER")
	dbPass := os.Getenv("PASS")
	dbName := os.Getenv("DBNAME")
	mongoDBURI := fmt.Sprintf("mongodb+srv://%s:%s@transformcluster-gelye.mongodb.net/%s?retryWrites=true&w=majority", dbUser, dbPass, dbName)

	clientOptions := options.Client().ApplyURI(mongoDBURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	transform := client.Database(dbName)
	clinicsCollection := transform.Collection("clinics")
	fashionCollection := transform.Collection("fashion")
	healthCollection := transform.Collection("health")
	historyCollection := transform.Collection("history")
	identityCollection := transform.Collection("identity")
	travelCollection := transform.Collection("travel")

	switch s {
	case "clinics":
		return clinicsCollection
	case "fashion":
		return fashionCollection
	case "health":
		return healthCollection
	case "history":
		return historyCollection
	case "identity":
		return identityCollection
	case "travel":
		return travelCollection
	}

	return clinicsCollection
}

// ErrorResponse: Struct for errors
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError: Helper function; creates error
func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())

	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode: http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}