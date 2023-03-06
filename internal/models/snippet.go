package models

import "time"

type Snippet struct {
	Id       int64
	Title    string `json:"title"`
	Content  string `json:"content"`
	Ctreated time.Time
	Expires  time.Time
}
