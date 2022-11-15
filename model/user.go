package model

import "time"

type UserProfile struct {
	Id        string    `db:"id" json:"id"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Gender    string    `db:"gender" json:"gender"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"pass" json:"password"`
	CreateAt  time.Time `db:"created_at" json:"-"`
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GenerateTokenResponse struct {
	Token string `json:"token"`
}
