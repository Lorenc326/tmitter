package tweet

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type repo struct {
	*sql.DB
}

func (r *repo) getById(id int) (Tweet, error) {
	row := r.DB.QueryRow("SELECT createdAt, username, text FROM tweets WHERE id=?", id)
	var res Tweet
	res.Id = id
	err := res.decode(row)
	return res, err
}

func (r *repo) getByUser() {
}
