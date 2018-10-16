package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
	query := "SELECT * FROM tasks"
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := TaskCollection{}

	for rows.Next() {
		task := Task{}
		rowErr := rows.Scan(&task.ID, &task.Name)

		if rowErr != nil {
			panic(rowErr)
		}

		result.Tasks = append(result.Tasks, task)
	}

	return result
}

func CreateTask(db *sql.DB, name string) (int64, error) {
	query := "INSERT INTO tasks(name) VALUES(?)"

	stmt, err := db.Prepare(query)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(name)

	if err != nil {
		panic(err)
	}

	return result.LastInsertId()
}

func UpdateTask(db *sql.DB, id int, name string) (int64, error) {
	return 0, nil
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	query := "DELETE FROM tasks WHERE id = ?"

	stmt, err := db.Prepare(query)

	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(id)

	if err != nil {
		panic(err)
	}

	return result.RowsAffected()
}
