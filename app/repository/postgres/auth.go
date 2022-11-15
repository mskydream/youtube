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
	tx := r.db.MustBegin()

	query := `INSERT INTO user_profile
				(first_name, last_name, gender, email, pass, created_at) 
				VALUES ($1, $2, $3, $4, $5, now())
				RETURNING id, first_name, last_name, gender, email, pass, created_at`

	err = tx.QueryRowx(query, input.FirstName, input.LastName, input.Gender, input.Email, input.Password).Scan(&res.Id, &res.FirstName, &res.LastName, &res.Gender, &res.Email, &res.Password, &res.CreateAt)
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = tx.Query(`INSERT INTO youtube_account(id, created_at) VALUES ($1, now())`, res.Id)
	if err != nil {
		tx.Rollback()
		return
	}

	return res, tx.Commit()
}

func (r *AuthPostgres) GetUser(input *model.SignIn) (user model.UserProfile, err error) {
	return user, r.db.Get(&user, "SELECT id, first_name, last_name, gender, email, pass, created_at from user_profile WHERE email = $1 AND pass = $2", input.Email, input.Password)
}
