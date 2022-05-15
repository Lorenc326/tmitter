package tweet

import (
	"database/sql"
	"encoding/json"
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

func GetTweetsByUser(c echo.Context) error {
	params := c.QueryParams()
	username := params.Get("username")
	if username == "" {
		return echo.ErrBadRequest
	}

	limit, _ := strconv.Atoi(params.Get("limit"))
	offset, _ := strconv.Atoi(params.Get("offset"))

	r := repo{db.Client}
	next, close, err := r.getByUser(username, limit, offset)
	if err == sql.ErrNoRows {
		return echo.ErrNotFound
	} else if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}
	defer close()

	// pipe data into response by chunks, so no intermediary results are stored in RAM
	enc := json.NewEncoder(c.Response())
	for {
		tweet := next()
		if tweet == nil {
			return nil
		}
		if err = enc.Encode(tweet); err != nil {
			log.Println(err)
			return echo.ErrInternalServerError
		}
		c.Response().Flush()
	}
}
