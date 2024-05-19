package handlers

import (
    "github.com/gorilla/mux"
    "github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitializeHandlers(r *mux.Router, database *sqlx.DB) {
    db = database
    r.HandleFunc("/users", CreateUserHandler).Methods("POST")
    r.HandleFunc("/users", ListUsersHandler).Methods("GET")
    r.HandleFunc("/users/{username}", DeleteUserHandler).Methods("DELETE")

    r.HandleFunc("/keys", CreateKeyHandler).Methods("POST")
    r.HandleFunc("/keys/{name}", DeleteKeyHandler).Methods("DELETE")

    r.HandleFunc("/items", CreateItemHandler).Methods("POST")
    r.HandleFunc("/items", ListItemsHandler).Methods("GET")
    r.HandleFunc("/items/{sector}/{name}", DeleteItemHandler).Methods("DELETE")
    r.HandleFunc("/items/{sector}/{name}", UpdateItemHandler).Methods("PUT")

    r.HandleFunc("/policies", CreatePolicyHandler).Methods("POST")
    r.HandleFunc("/policies", ListPoliciesHandler).Methods("GET")
    r.HandleFunc("/policies/{name}", DeletePolicyHandler).Methods("DELETE")
}
