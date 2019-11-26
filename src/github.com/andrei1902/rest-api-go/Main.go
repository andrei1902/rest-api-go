package main

import (
	"github.com/gorilla/mux"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User struct
type User struct {
	Id   primitive.ObjectID `json:"_id, omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson: "name,omitempty"`
	Email string             `json:"email,omitempty" bson: "email,omitempty"`
}

var client *mongo.Client

func CreateUserEndPoint(response http.ResponseWriter, request *http.ReadRequest) {
	response.Header().Add("content-type", "application/json")
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("AndreiCluster").Collection("users")
	ctx, _ := contex.WithTimeout(context.Background(), 10*time.Second)
	result,  := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

func main() {
	fmt.Println("Hello from GO!")
	ctx, _ := contex.WithTimeout(context.Background(), 10*time.Second)
	client, + = mongo.Connect(ctx, "mongodb+srv://andrei123:Zsxo19qk7KUnEqwo@andreicluster-dkx3z.mongodb.net/test?retryWrites=true^&w=majority")

	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUserEndPoint).Methods("POST")
	http.ListenAndServe(":12345", router)
}
