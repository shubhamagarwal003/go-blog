package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string
}

type SessionToken struct {
	Id    int
	Token string
	Uid   int
}
