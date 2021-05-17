package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	workshop_2 "github.com/zemags/go_workshop_2/store"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userID int, list workshop_2.TodoList) (int, error) {
	// Run transaction
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("insert into %s (title, description) values ($1, $2) returning id", todoListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserListQuery := fmt.Sprintf("insert into %s (user_id, list_id) values ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUserListQuery, userID, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userID int) ([]workshop_2.TodoList, error) {
	var lists []workshop_2.TodoList

	query := fmt.Sprintf(
		`select tl.id, tl.title, tl.description from %s tl
		inner join %s ul
		on tl.id  = ul.list_id
		where ul.user_id = $1;`,
		todoListTable, usersListsTable,
	)

	err := r.db.Select(&lists, query, userID)
	return lists, err
}

func (r *TodoListPostgres) GetByID(userID, listID int) (workshop_2.TodoList, error) {
	var list workshop_2.TodoList

	query := fmt.Sprintf(
		`select tl.id, tl.title, tl.description from %s tl
		inner join %s ul
		on tl.id  = ul.list_id
		where ul.user_id = $1
		and ul.list_id = $2;`,
		todoListTable, usersListsTable,
	)

	err := r.db.Get(&list, query, userID, listID)
	return list, err
}

func (r *TodoListPostgres) Delete(userID, listID int) error {
	query := fmt.Sprintf("delete from %s tl using %s ul where tl.id = ul.list_id and ul.user_id=$1 and ul.list_id=$2",
		todoListTable, usersListsTable)

	_, err := r.db.Exec(query, userID, listID)
	return err
}

func (r *TodoListPostgres) Update(userID, listID int, input workshop_2.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *input.Description)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("update %s tl set %s from %s ul where tl.id = ul.list_id and ul.list_id=$%d and ul.user_id=$%d",
		todoListTable, setQuery, usersListsTable, argID, argID+1)
	args = append(args, listID, userID)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
