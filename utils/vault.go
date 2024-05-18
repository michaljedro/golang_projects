package utils

import (
    "log"
    "os"

    "github.com/hashicorp/vault/api"
)

var VaultClient *api.Client

func InitVault() {
    config := &api.Config{
        Address: "http://localhost:8200",
    }

    client, err := api.NewClient(config)
    if err != nil {
        log.Fatalf("Unable to initialize Vault client: %v", err)
    }

    client.SetToken(os.Getenv("VAULT_TOKEN"))
    VaultClient = client
}
