package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/stevegood/ike/models"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

func PostTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var task models.Task

		c.Bind(&task)

		id, err := models.CreateTask(db, task.Name)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, H{
			"created": id,
		})
	}
}

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, H{
			"updated": id,
		})
	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		_, err := models.DeleteTask(db, id)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}
