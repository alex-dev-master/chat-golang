package repository

import (
	"fmt"
	"github.com/alex-dev-master/chat-golang/pkg/model"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"time"
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

func (r *Auth) GetUser(email, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=? AND password=?", "users")
	err := r.db.Get(&user, query, email, password)

	return user, err
}

func (r *Auth) GetByRefreshToken(refreshToken string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE token=?", "users")
	err := r.db.Get(&user, query, refreshToken)

	return user, err
}

func (r *Auth) UpdateRefreshToken(user model.User, refreshToken string) error {
	refreshTokenTTLStr := viper.GetString("auth.refreshTokenTTL")
	refreshTokenTTL, _ := time.ParseDuration(refreshTokenTTLStr)
	var now time.Time
	now = time.Now()
	timeExpired := now.Add(time.Minute + refreshTokenTTL).Unix()

	query := fmt.Sprintf("UPDATE %s SET token=?, expired_token=? WHERE id=?", "users")
	_, err := r.db.Exec(query, refreshToken, timeExpired, user.Id)

	return err
}