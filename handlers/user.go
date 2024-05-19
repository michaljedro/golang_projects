package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "myapp/models"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Use the global db variable here
    if _, err := db.NamedExec(`INSERT INTO users (username, email) VALUES (:username, :email)`, &user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
    var users []models.User
    if err := db.Select(&users, "SELECT username, email FROM users"); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    username := params["username"]

    if _, err := db.Exec("DELETE FROM users WHERE username = ?", username); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
