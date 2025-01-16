package models

type User struct {
	ID       int    `gorm:"primary_key"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Email    string `gorm:"unique" json:"email"`
}

func (u *User) TableName() string {
	return "users"
}
