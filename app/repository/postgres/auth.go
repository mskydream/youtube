package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/mskydream/youtube/model"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) SignUp(input *model.UserProfile) (res model.UserProfile, err error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return
	}
	defer tx.Rollback()

	query := `INSERT INTO user_profile
				(first_name, last_name, gender, email, pass, created_at) 
				values ($1, $2, $3, $4, $5, now())
				RETURNING id, first_name, last_name, gender, email, pass, created_at`

	row := r.db.QueryRow(query, input.FistName, input.LastName, input.Gender, input.Email, input.Password)
	if err = row.Scan(&res.ID, &res.FistName, &res.LastName, &res.Gender, &res.Email, &res.Password, &res.CreateAt); err != nil {
		return
	}

	return res, nil
}

func (r *AuthPostgres) GetUser(input *model.SignIn) (user model.UserProfile, err error) {
	return user, r.db.Get(&user, "select id, first_name, last_name, gender, email, pass, created_at from user_profile where email = $1 AND pass = $2", input.Email, input.Password)
}
