package data

import (
	"errors"
	"time"
)

type Room struct {
	Id        int
	UUId      string
	Name      string
	UserId    int
	Token     string
	CreatedAt time.Time
}

func RoomByUUid(uuid string) (r Room, err error) {
	q := `SELECT ID, UUID, NAME, USER_ID, TOKEN, CREATED_AT FROM ROOMS WHERE UUID = $1`
	err = db.QueryRow(q, uuid).Scan(&r.Id, &r.UUId, &r.Name, &r.UserId, &r.Token, &r.CreatedAt)
	if err != nil || r.Id < 1 {
		err = errors.New("Not found room by uuid: "+uuid)
	}
	return
}

func AllRooms() (rs []Room, err error) {
	q := `SELECT ID, UUID, NAME, USER_ID, TOKEN, CREATED_AT FROM ROOMS;`
	rows, err := db.Query(q)
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

func (r *Room) Update() (err error) {
	q := `UPDATE ROOMS SET NAME = $1, TOKEN = $2 WHERE ID = $3;`
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	defer func() { _ = stmt.Close() }()
	_, err = stmt.Exec(r.Name, r.Token, r.Id)
	return
}

func (r *Room) Delete() (err error) {
	q1 := `DELETE FROM JOIN_ROOM WHERE ROOM_ID = $1;`
	q2 := `DELETE FROM ROOMS WHERE ID = $1;`
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() { _ = tx.Commit() }()
	_, err = tx.Exec(q1, r.Id)
	if err != nil {
		return
	}
	_, err = tx.Exec(q2, r.Id)
	return
}

func (r *Room) Creator() (u User, err error) {
	q := `SELECT ID, UUID, NAME, EMAIL, PASSWORD, MAX_ROLE, CREATED_AT FROM USERS WHERE ID = $1;`
	err = db.QueryRow(q, r.UserId).Scan(&u.Id, &u.UUId, &u.Name, &u.Email, &u.Password, &u.MaxRole, &u.CreatedAt)
	return
}

func (r *Room) Users() (us []User, err error) {
	q := `SELECT U.ID, U.UUID, NAME, EMAIL, PASSWORD, MAX_ROLE, CREATED_AT FROM USERS U, JOIN_ROOM J WHERE U.ID = J.USER_ID AND J.ROOM_ID = $1`
	rows, err := db.Query(q, r.UserId)
	if nil == rows || err != nil {
		return
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var u User
		if err = rows.Scan(&u.Id, &u.UUId, &u.Name, &u.Email, &u.Password, &u.MaxRole, &u.CreatedAt); err != nil {
			return
		}
		us = append(us, u)
	}
	return
}

func (r *Room) CreatedTime() string {
	return r.CreatedAt.Format("2006-01-02 03:04:05")
}
