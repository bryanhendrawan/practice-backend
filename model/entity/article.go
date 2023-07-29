package entity

import "time"

type Article struct {
	ID      int       `json:"id" gorm:"primaryKey"`
	Author  string    `json:"author"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}
