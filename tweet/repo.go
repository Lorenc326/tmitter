package tweet

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type repo struct {
	*sql.DB
}

func (r *repo) getById(id int) (Tweet, error) {
	row := r.QueryRow("SELECT id, createdAt, username, text FROM tweets WHERE id=?", id)
	var res Tweet
	err := res.decode(row)
	return res, err
}

func (r *repo) getByUser(username string, limit int, offset int) (*[]UserPopulated, error) {
	rows, err := r.Query(`
SELECT tweets.id, tweets.createdAt, tweets.username, tweets.text, users.firstName, users.lastName
FROM tweets 
LEFT JOIN users USING (username)
WHERE tweets.username = ?
LIMIT ?, ?;
`, username, offset, limit)
	defer rows.Close()

	res := []UserPopulated{}
	if err != nil {
		return &res, err
	}

	for rows.Next() {
		var tweet UserPopulated
		tweet.decode(rows)
		log.Println(tweet)
		if err != nil {
			return &res, err
		}
		res = append(res, tweet)
	}

	return &res, nil
}
