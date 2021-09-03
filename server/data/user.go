package data

import (
	"errors"
	"github.com/duruyao/gochat/server/util"
	"time"
)

type Role int

const (
	Owner Role = 0
	Admin Role = 2
	Guest Role = 4
)

type User struct {
	Id        int
	UUId      string
	Name      string
	Email     string
	Password  string
	MaxRole   Role
	CreatedAt time.Time
}

func UserByUUId(uuid string) (u User, err error) {
	q := `SELECT ID, UUID, NAME, EMAIL, PASSWORD, MAX_ROLE, CREATED_AT FROM USERS WHERE UUID = $1;`
	err = db.QueryRow(q, uuid).Scan(&u.Id, &u.UUId, &u.Name, &u.Email, &u.Password, &u.MaxRole, &u.CreatedAt)
	if err != nil || u.Id < 1 {
		err = errors.New("not found")
	}
	return
}

func AllUsers() (us []User, err error) {
	q := `SELECT ID, UUID, NAME, EMAIL, PASSWORD, MAX_ROLE, CREATED_AT FROM USERS;`
	rows, err := db.Query(q)
	if nil == rows || err != nil {
		return
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		u := User{}
		if err = rows.Scan(&u.Id, &u.UUId, &u.Name, &u.Email, &u.Password, &u.MaxRole, &u.CreatedAt); err != nil {
			return
		}
		us = append(us, u)
	}
	return
}

func (u *User) Create() (err error) {
	q1 := `INSERT INTO USERS (UUID, NAME, EMAIL, PASSWORD, MAX_ROLE, CREATED_AT) VALUES ($1, $2, $3, $4, $5, $6);`
	q2 := `SELECT LAST_INSERT_ROWID();`
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() { _ = tx.Commit() }()
	u.UUId = util.CreateUUId()
	u.CreatedAt = time.Now().Local()
	_, err = tx.Exec(q1, u.UUId, u.Name, u.Email, util.Encrypt(u.Password), u.MaxRole, u.CreatedAt)
	if err != nil {
		return
	}
	err = tx.QueryRow(q2).Scan(&u.Id)
	return
}

func (u *User) Update() (err error) {
	q := `UPDATE USERS SET NAME = $1, EMAIL = $2, PASSWORD = $3, MAX_ROLE = $4 WHERE ID = $5;`
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer func() { _ = stmt.Close() }()
	_, err = stmt.Exec(u.Name, u.Email, u.Password, u.MaxRole, u.Id)
	return
}

func (u *User) Delete() (err error) {
	q1 := `DELETE FROM JOIN_ROOM WHERE USER_ID = $1;`
	q2 := `DELETE FROM ROOMS WHERE USER_ID = $1;`
	q3 := `DELETE FROM USERS WHERE ID = $1;`
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() { _ = tx.Commit() }()
	_, err = tx.Exec(q1, u.Id)
	if err != nil {
		return
	}
	_, err = tx.Exec(q2, u.Id)
	if err != nil {
		return
	}
	_, err = tx.Exec(q3, u.Id)
	return
}

func (u *User) CreateSession() (s Session, err error) {
	q1 := `INSERT INTO SESSIONS (UUID, USER_ID, CREATED_AT) VALUES ($1, $2, $3);`
	q2 := `SELECT LAST_INSERT_ROWID();`
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() { _ = tx.Commit() }()
	s = Session{
		UUId:      util.CreateUUId(),
		UserId:    u.Id,
		CreatedAt: time.Now().Local(),
	}
	_, err = tx.Exec(q1, s.UUId, s.UserId, s.CreatedAt)
	if err != nil {
		return
	}
	err = tx.QueryRow(q2).Scan(&s.Id)
	return
}

func (u *User) Session() (s Session, err error) {
	q := `SELECT ID, UUID, USER_ID, CREATED_AT FROM SESSIONS WHERE USER_ID = $1;`
	err = db.QueryRow(q, u.Id).Scan(&s.Id, &s.UUId, &s.UserId, &s.CreatedAt)
	return
}

func (u *User) CreateRoom(name string, token string) (r Room, err error) {
	q1 := `INSERT INTO ROOMS (UUID, NAME, USER_ID, TOKEN) VALUES ($1, $2, $3, $4);`
	q2 := `SELECT LAST_INSERT_ROWID();`
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() { _ = tx.Commit() }()
	r = Room{
		UUId:      util.CreateUUId(),
		Name:      name,
		UserId:    u.Id,
		Token:     token,
		CreatedAt: time.Now().Local(),
	}
	_, err = tx.Exec(q1, r.UUId, r.Name, r.UserId, r.Token)
	if err != nil {
		return
	}
	err = tx.QueryRow(q2).Scan(&r.Id)
	return
}

func (u *User) CreatedRooms() (rs []Room, err error) {
	q := `SELECT ID, UUID, NAME, USER_ID, TOKEN, CREATED_AT FROM ROOMS WHERE USER_ID = $1;`
	rows, err := db.Query(q, u.Id)
	if nil == rows || err != nil {
		return
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		r := Room{}
		if err = rows.Scan(&r.Id, &r.UUId, &r.Name, &r.UserId, &r.Token, &r.CreatedAt); err != nil {
			return
		}
		rs = append(rs, r)
	}
	return
}

func (u *User) JoinRoom(r Room) (err error) {
	q := `INSERT INTO JOIN_ROOM (UUID, ROOM_ID, USER_ID) VALUES ($1, $2, $3);`
	_, err = db.Exec(q, util.CreateUUId(), r.Id, u.Id)
	return
}

func (u *User) JoinedRooms() (rs []Room, err error) {
	q := `SELECT R.ID, R.UUID, NAME, TOKEN, CREATED_AT FROM ROOMS R, JOIN_ROOM J WHERE R.ID = J.ROOM_ID AND J.USER_ID = $1;`
	rows, err := db.Query(q, u.Id)
	if nil == rows || err != nil {
		return
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var r Room
		if err = rows.Scan(&r.Id, &r.UUId, &r.Name, &r.Token, &r.CreatedAt); err != nil {
			return
		}
		rs = append(rs, r)
	}
	return
}

func (u *User) LeaveRoom(r Room) (err error) {
	q := `DELETE FROM JOIN_ROOM WHERE ROOM_ID = $1;`
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer func() { _ = stmt.Close() }()
	_, err = stmt.Exec(r.Id)
	return
}

func (u *User) CreatedTime() string {
	return u.CreatedAt.Format("2006-01-02 03:04:05")
}
