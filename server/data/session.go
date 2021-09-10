package data

import (
	"fmt"
	"time"
)

type Session struct {
	Id        int       `db:"ID" json:"id" form:"id" uri:"id"`
	UUId      string    `db:"UUID" json:"uuid" form:"uuid"`
	UserId    int       `db:"USER_ID" json:"user_id" form:"user_id"`
	CreatedAt time.Time `db:"CREATED_AT" json:"created_at"`
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
