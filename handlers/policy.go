package handlers

import (
    "encoding/json"
    "net/http"
    "myapp/models"
    "myapp/utils"
)

func CreatePolicy(w http.ResponseWriter, r *http.Request) {
    var policy models.Policy
    if err := json.NewDecoder(r.Body).Decode(&policy); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if _, err := utils.DB.NamedExec(`INSERT INTO policies (name, definition) VALUES (:name, :definition)`, &policy); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}
