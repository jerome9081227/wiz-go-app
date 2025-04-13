package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    data, err := os.ReadFile("wizexercise.txt")
    if err != nil {
        log.Fatalf("wizexercise.txt missing or unreadable: %v", err)
    }
    log.Printf("✔ wizexercise.txt content: %s", string(data))


    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI not set")
    }

    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    defer client.Disconnect(context.TODO())

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "WIZ Go App is Live 🚀")
    })

    http.ListenAndServe(":8080", nil)
}
