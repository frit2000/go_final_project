package servicetask

import "database/sql"

type TaskStore struct {
	db *sql.DB
}

type Task struct {
	Id      string `json:"id,omitempty"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat,omitempty"`
}
