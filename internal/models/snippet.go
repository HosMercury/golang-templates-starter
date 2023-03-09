package models

import "time"

type Snippet struct {
	Id      int64
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Created time.Time
	Expires time.Time
	Version int32
}
