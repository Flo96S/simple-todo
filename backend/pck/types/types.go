package types

import "time"

type TodoItem struct {
	Id          string
	Title       string
	Description string
	Created     time.Time
	IsDone      bool
	Image       string
	Color       string
}

type TodoItemShort struct {
	Id    string
	Title string
	Color string
}
