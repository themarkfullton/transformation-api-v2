package main

import(
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/themarkfullton/transformation-api-v2/helper"
	"github.com/themarkfullton/transformation-api-v2/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getAllClinics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var clinics []models.Clinic

	collection := helper.ConnectDB("clinics")

	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var clinic models.Clinic
		err := cursor.Decode(&clinic)

		if err != nil {
			log.Fatal(err)
		}

		clinics = append(clinics, clinic)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(clinics)
}

func main(){
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./static")))

	port := os.Getenv("PORT")

	if err := http.ListenAndServe(":"+ port, r); err != nil {

		log.Fatal(err)
	}
}