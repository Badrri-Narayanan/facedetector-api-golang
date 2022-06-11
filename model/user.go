package model

import "time"

type User struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Entries int       `json:"entries"`
	Joined  time.Time `json:"joined"`
	Age     int       `json:"age"`
	Pet     string    `json:"pet"`
}
