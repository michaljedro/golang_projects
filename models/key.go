package models

type Key struct {
    Name     string `json:"name"`
    Secret   string `json:"secret" gorm:"-"`
    Username string `json:"username"`
}
