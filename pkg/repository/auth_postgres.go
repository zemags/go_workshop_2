package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	workshop_2 "github.com/zemags/go_workshop_2/store"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user workshop_2.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (name, username, password_hash) values ($1, $2, $3) returning id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (workshop_2.User, error) {
	var user workshop_2.User
	query := fmt.Sprintf(
		"select id from %s where username=$1 and password_hash=$2",
		userTable,
	)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
