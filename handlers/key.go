package handlers

import (
    "encoding/json"
    "net/http"
    "myapp/models"
    "myapp/utils"
)

func CreateKey(w http.ResponseWriter, r *http.Request) {
    var key models.Key
    if err := json.NewDecoder(r.Body).Decode(&key); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Zapisz secret w Vault
    secretPath := "secret/data/keys/" + key.Name
    secretData := map[string]interface{}{
        "data": map[string]interface{}{
            "secret": key.Secret,
        },
    }
    _, err := utils.VaultClient.Logical().Write(secretPath, secretData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    key.Secret = ""
    if _, err := utils.DB.NamedExec(`INSERT INTO keys (name, username) VALUES (:name, :username)`, &key); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}
