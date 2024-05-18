package models

type User struct {
    Username string `json:"username" gorm:"primaryKey"`
    Email    string `json:"email"`
    Policies []Policy `json:"policies" gorm:"many2many:user_policies;"`
}
