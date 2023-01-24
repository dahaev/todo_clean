package repository

import (
	"fmt"

	"github.com/dahaev/todo.git"
	"github.com/jmoiron/sqlx"
)

type TodoListPorstgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPorstgres {
	return &TodoListPorstgres{
		db: db,
	}
}

func (r *TodoListPorstgres) Create(userID int, list todo.Todolist) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int

	createListQuery := fmt.Sprintf("INSERT INTO %s(title, description) VALUES($1, $2)", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersList := fmt.Sprintf("INSERT INTO %s(user_id, list_id) VALUES($1, $2)", userListsTable)
	_, err = tx.Exec(createUsersList, userID, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()

}
