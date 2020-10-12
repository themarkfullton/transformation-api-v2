package transformation_api_v2

import(
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/themarkfullton/transformation-api-read-only/helper"
	"github.com/themarkfullton/transformation-api-read-only/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main(){
	r := mux.NewRouter()

	r.Handle("/", http.FileServer())

	port := os.Getenv("PORT")

	if err := http.ListenAndServe(":" + port, r); err != nil {
		log.Fatal(err)
	}
}