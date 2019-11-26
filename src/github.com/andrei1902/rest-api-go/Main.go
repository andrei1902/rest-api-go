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
	ID   primitive.ObjectID `json:"_id, omitempty" bson:"_id,omitempty"`
	Name string             `bson:"Name,omitempty" bson: "Name,omitempty"`
}

var client *mongo.Client

func main() {
	fmt.Println("Hello from GO!")
	ctx, _ := contex.WithTimeout(context.Background(), 10*time.Second)
	client, + = mongo.Connect(ctx, "mongodb+srv://andrei123:Zsxo19qk7KUnEqwo@andreicluster-dkx3z.mongodb.net/test?retryWrites=true^&w=majority")

	router := mux.NewRouter()
	http.ListenAndServe(":12345", router)
}
