package data

import (
	"errors"
	"time"
)

type Session struct {
	Id        int
	UUId      string
	UserId    int
	CreatedAt time.Time
}

func SessionByUUid(uuid string) (s Session, err error) {
	q := `SELECT ID, UUID, USER_ID, CREATED_AT FROM SESSIONS WHERE UUID = $1`
	err = db.QueryRow(q, uuid).Scan(&s.Id, &s.UUId, &s.UserId, &s.CreatedAt)
	if err != nil || s.Id < 1 {
		err = errors.New("Not found session by uuid: " + uuid)
	}
	return
}

func AllSessions() (ss []Session, err error) {
	q := `SELECT ID, UUID, USER_ID, CREATED_AT FROM SESSIONS;`
	rows, err := db.Query(q)
	if nil == rows || err != nil {
		return
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		s := Session{}
		if err = rows.Scan(&s.Id, &s.UUId, &s.UserId, &s.CreatedAt); err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

//
//func (s *Session) CheckUUId() (valid bool, err error) {
//	q := `SELECT ID, UUID, USER_ID, CREATED_AT FROM SESSIONS WHERE UUID = $1;`
//	err = db.QueryRow(q, s.UUId).Scan(&s.Id, &s.UUId, &s.UserId, &s.CreatedAt)
//	if err != nil {
//		return
//	}
//	valid = s.Id > 0
//	return
//}

func (s *Session) Delete() (err error) {
	q := `DELETE FROM SESSIONS WHERE ID = $1;`
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer func() { _ = stmt.Close() }()
	_, err = stmt.Exec(s.Id)
	return
}

func (s *Session) User() (u User, err error) {
	q := `SELECT ID, UUID, NAME, EMAIL, PASSWORD, MAX_ROLE, CREATED_AT FROM USERS WHERE ID = $1;`
	err = db.QueryRow(q, s.UserId).Scan(&u.Id, &u.UUId, &u.Name, &u.Email, &u.Password, &u.MaxRole, &u.CreatedAt)
	return
}

func (s *Session) CreatedTime() string {
	return s.CreatedAt.Format("2006-01-02 03:04:05")
}
