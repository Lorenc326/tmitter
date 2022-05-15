package tweet

import (
	"database/sql"
	"github.com/Lorenc326/tmitter/db"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
)

func GetTweet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	r := repo{db.Client}
	tweet, err := r.getById(id)
	if err == sql.ErrNoRows {
		return echo.ErrNotFound
	} else if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	return c.JSON(200, tweet)
}
