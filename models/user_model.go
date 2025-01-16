package models

type User struct {
	ID    int    `gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `gorm:"unique" json:"email"`
}

func (u *User) TableName() string {
	return "users"
}
