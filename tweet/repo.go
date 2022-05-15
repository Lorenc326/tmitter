package tweet

import (
	"database/sql"
	"github.com/Lorenc326/tmitter/db"
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

func (r *repo) getByUser(username string, limit int, offset int) (db.Iterate, db.CloseIterator, error) {
	rows, err := r.Query(`
SELECT tweets.id, tweets.createdAt, tweets.username, tweets.text, users.firstName, users.lastName
FROM tweets 
LEFT JOIN users USING (username)
WHERE tweets.username = ?
ORDER BY tweets.createdAt ASC
LIMIT ?, ?;
`, username, offset, limit)

	if err != nil {
		return nil, nil, err
	}

	iterator := func() interface{} {
		if rows.Next() {
			var tweet UserPopulated
			tweet.decode(rows)
			if err != nil {
				log.Println(err)
				return nil
			}
			return tweet
		}
		return nil
	}
	closer := func() {
		rows.Close()
	}
	return iterator, closer, nil
}
