package main

import (
    "log"
    "net/http"
    "myapp/handlers"
    "myapp/utils"

    "github.com/gorilla/mux"
)

func main() {
    utils.InitDB()
    utils.InitVault()

    r := mux.NewRouter()

    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/keys", handlers.CreateKey).Methods("POST")
    r.HandleFunc("/items", handlers.CreateItem).Methods("POST")
    r.HandleFunc("/policies", handlers.CreatePolicy).Methods("POST")

    log.Fatal(http.ListenAndServe(":8080", r))
}
