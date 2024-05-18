package handlers

import (
    "encoding/json"
    "net/http"
    "myapp/models"
    "myapp/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if _, err := utils.DB.NamedExec(`INSERT INTO users (username, email) VALUES (:username, :email)`, &user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}
