package main

import (
	"database/sql"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/stevegood/ike/handlers"
)

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	q := `
		CREATE TABLE IF NOT EXISTS tasks(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL
		)
	`

	_, err := db.Exec(q)

	// exit if something goes wrong with the sql statement
	if err != nil {
		panic(err)
	}
}

func main() {
	db := initDB("ike.db")
	migrate(db)

	// create a new instance of Echo
	e := echo.New()

	e.Static("/static", "client/build/static")
	e.File("/", "client/build/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.POST("/tasks", handlers.PostTask(db))
	e.PUT("/tasks/:id", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	// start as a web server
	e.Logger.Fatal(e.Start(":8000"))
}
