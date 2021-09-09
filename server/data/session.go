package data

import (
	"fmt"
	"time"
)

type Session struct {
	Id        int       `db:"ID"`
	UUId      string    `db:"UUID"`
	UserId    int       `db:"USER_ID"`
	CreatedAt time.Time `db:"CREATED_AT"`
}

func SessionByUniqueKey(key string, value interface{}) (s Session, err error) {
	q := fmt.Sprintf("SELECT * FROM SESSIONS WHERE %s = $2;", key)
	err = db.Get(&s, q, value)
	return
}

func Sessions(limit int) (ss []Session, err error) {
	q := `SELECT * FROM SESSIONS LIMIT $1;`
	err = db.Select(&ss, q, limit)
	return
}

func (s *Session) Delete() (err error) {
	q := `DELETE FROM SESSIONS WHERE ID = $1;`
	_, err = db.Exec(q, s.Id)
	return
}

func (s *Session) User() (u User, err error) {
	q := `SELECT * FROM USERS WHERE ID = $1;`
	err = db.Get(&u, q, s.UserId)
	return
}
