package repository

import (
	"fmt"
	"github.com/alex-dev-master/chat-golang/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}


func (r *Auth) CreateUser(user model.User) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, email, password, token) values (?,?,?,?,?)", "users")

	row, err := r.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, user.Token)
	if err != nil {
		return 0, err
	}
	id, err = row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

