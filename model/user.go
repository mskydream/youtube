package model

import "time"

type UserProfile struct {
	ID       int64     `json:"id"`
	FistName string    `json:"first_name"`
	LastName string    `json:"last_name"`
	Gender   string    `json:"gender"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"created_at"`
}
