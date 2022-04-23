package controllers

type User struct {
	Username string `json:"username"`
	Number   string `json:"number"`
	IsAdmin  int8   `json:"is_admin"`
	CreateNews int8 `json:"create_news"`
}

func NewUser() *User {
	return &User{}
}