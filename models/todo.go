package models

type Todo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Completed bool  `json:"completed"`
	UserID   int    `json:"user_id"`
}
