package repository

import (
	"fmt"

	"github.com/Dorkhan1/todo-app"
	"github.com/jmoiron/sqlx"
)

type TodoListPostrgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostrgres {
	return &TodoListPostrgres{db: db}
}

func (r *TodoListPostrgres) Create(userId int, list todo.ToDoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	CreateListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(CreateListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	CreateUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(CreateUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()

}

func (r *TodoListPostrgres) GetAll(userId int) ([]todo.ToDoList, error) {
	var lists []todo.ToDoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err

}

func (r *TodoListPostrgres) GetById(userId, listId int) (todo.ToDoList, error) {
	var list todo.ToDoList
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
								 INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err


}
