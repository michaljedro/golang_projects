package handlers

import (
    "encoding/json"
    "net/http"
    "myapp/models"
    "myapp/utils"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
    var item models.Item
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if _, err := utils.DB.NamedExec(`INSERT INTO items (sector, name, description) VALUES (:sector, :name, :description)`, &item); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}
