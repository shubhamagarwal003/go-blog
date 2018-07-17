package models

type Blog struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Uid      int64  `json:"uid"`
	Username string `json:"user"`
}
