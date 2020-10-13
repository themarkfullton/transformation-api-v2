package main

import(
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/themarkfullton/transformation-api-v2/helper"
	"github.com/themarkfullton/transformation-api-v2/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ==============================/ CLINICS / ======================>
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

// ==============================/ RESOURCES / ======================>

func getAllResources(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	fmt.Printf("%v", r)
}

// ==============================/ NEWS / ======================>

func getAllNews(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var announcements []models.Announcement

	collection := helper.ConnectDB("news")

	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var announcement models.Announcement
		err := cursor.Decode(&announcement)

		if err != nil {
			log.Fatal(err)
		}

		announcements = append(announcements, announcement)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(announcements)
}

// ==============================/ MAIN / ======================>
func main(){
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./static")))

	r.HandleFunc("/api/clinics", getAllClinics).Methods("GET")
	r.HandleFunc("/api/fashion", getAllResources).Methods("GET")
	r.HandleFunc("/api/health", getAllResources).Methods("GET")
	r.HandleFunc("/api/history", getAllResources).Methods("GET")
	r.HandleFunc("/api/identity", getAllResources).Methods("GET")
	r.HandleFunc("/api/travel", getAllResources).Methods("GET")
	r.HandleFunc("/api/news", getAllNews).Methods("GET")


	port := os.Getenv("PORT")

	if err := http.ListenAndServe(":"+ port, r); err != nil {

		log.Fatal(err)
	}
}