package tweet

import "database/sql"

type Tweet struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"createdAt"`
	Username  string `json:"username"`
	Text      string `json:"text"`
}

func (tweet *Tweet) decode(row *sql.Row) error {
	return row.Scan(&tweet.Id, &tweet.CreatedAt, &tweet.Username, &tweet.Text)
}
