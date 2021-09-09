package data

import (
	"fmt"
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
	Id        int       `db:"ID"`
	UUId      string    `db:"UUID"`
	Name      string    `db:"NAME"`
	Password  string    `db:"PASSWORD"`
	MaxRole   Role      `db:"MAX_ROLE"`
	CreatedAt time.Time `db:"CREATED_AT"`
}

func UserByUniqueKey(key string, value interface{}) (u User, err error) {
	q := fmt.Sprintf("SELECT * FROM USERS WHERE %s = $2;", key)
	err = db.Get(&u, q, value)
	return
}

func Users(limit int) (us []User, err error) {
	q := `SELECT * FROM USERS LIMIT $1;`
	err = db.Select(&us, q, limit)
	return
}

func (u *User) Create() (err error) {
	q1 := `INSERT INTO USERS (UUID, NAME, PASSWORD, MAX_ROLE, CREATED_AT) VALUES ($1, $2, $3, $4, $5);`
	q2 := `SELECT LAST_INSERT_ROWID();`
	tx, err := db.Beginx()
	if err != nil {
		return
	}
	defer func() { _ = tx.Commit() }()
	u.UUId = util.CreateUUId()
	u.CreatedAt = time.Now().Local()
	if _, err = tx.Exec(q1, u.UUId, u.Name, util.Encrypt(u.Password), u.MaxRole, u.CreatedAt); err != nil {
		return
	}
	err = tx.Get(&u.Id, q2)
	return
}

func (u *User) Update() (err error) {
	q := `UPDATE USERS SET NAME = $1, PASSWORD = $3, MAX_ROLE = $4 WHERE ID = $5;`
	_, err = db.Exec(q, u.Name, u.Password, u.MaxRole, u.Id)
	return
}

func (u *User) Delete() (err error) {
	q1 := `DELETE FROM MEMBERS WHERE USER_ID = $1;`
	q2 := `DELETE FROM GROUPS WHERE ADMIN_ID = $1;`
	q3 := `DELETE FROM SESSIONS WHERE USER_ID = $1;`
	q4 := `DELETE FROM USERS WHERE ID = $1;`
	tx, err := db.Beginx()
	if err != nil {
		return
	}
	defer func() { _ = tx.Commit() }()
	if _, err = tx.Exec(q1, u.Id); err != nil {
		return
	}
	if _, err = tx.Exec(q2, u.Id); err != nil {
		return
	}
	if _, err = tx.Exec(q3, u.Id); err != nil {
		return
	}
	_, err = tx.Exec(q4, u.Id)
	return
}

func (u *User) CreateSession() (s Session, err error) {
	q1 := `INSERT INTO SESSIONS (UUID, USER_ID, CREATED_AT) VALUES ($1, $2, $3);`
	q2 := `SELECT LAST_INSERT_ROWID();`
	tx, err := db.Beginx()
	if err != nil {
		return
	}
	defer func() { _ = tx.Commit() }()
	s = Session{
		UUId:      util.CreateUUId(),
		UserId:    u.Id,
		CreatedAt: time.Now().Local(),
	}
	if _, err = tx.Exec(q1, s.UUId, s.UserId, s.CreatedAt); err != nil {
		return
	}
	err = tx.Get(&s.Id, q2)
	return
}

func (u *User) Session() (s Session, err error) {
	q := `SELECT * FROM SESSIONS WHERE USER_ID = $1;`
	err = db.Get(&s, q, u.Id)
	return
}

func (u *User) CreateGroup(name string, token string) (g Group, err error) {
	q1 := `INSERT INTO GROUPS (UUID, NAME, ADMIN_ID, TOKEN) VALUES ($1, $2, $3, $4);`
	q2 := `SELECT LAST_INSERT_ROWID();`
	q3 := `INSERT INTO MEMBERS (UUID, GROUP_ID, USER_ID) VALUES ($1, $2, $3);`
	tx, err := db.Beginx()
	if err != nil {
		return
	}
	defer func() { _ = tx.Commit() }()
	g = Group{
		UUId:      util.CreateUUId(),
		Name:      name,
		AdminId:   u.Id,
		Token:     token,
		CreatedAt: time.Now().Local(),
	}
	if _, err = tx.Exec(q1, g.UUId, g.Name, g.AdminId, g.Token); err != nil {
		return
	}
	if err = tx.Get(&g.Id, q2); err != nil {
		return
	}
	_, err = tx.Exec(q3, util.CreateUUId(), g.Id, u.Id)
	return
}

func (u *User) CreatedGroups() (gs []Group, err error) {
	q := `SELECT * FROM GROUPS WHERE ADMIN_ID = $1;`
	err = db.Select(&gs, q, u.Id)
	return
}

func (u *User) JoinGroup(g Group) (err error) {
	q := `INSERT INTO MEMBERS (UUID, GROUP_ID, USER_ID) VALUES ($1, $2, $3);`
	_, err = db.Exec(q, util.CreateUUId(), g.Id, u.Id)
	return
}

func (u *User) JoinedGroups() (gs []Group, err error) {
	q := `SELECT G.* FROM GROUPS G, MEMBERS M WHERE G.ID = M.GROUP_ID AND M.USER_ID = $1;`
	err = db.Select(&gs, q, u.Id)
	return
}

func (u *User) LeaveGroup(g Group) (err error) {
	q := `DELETE FROM MEMBERS WHERE GROUP_ID = $1;`
	_, err = db.Exec(q, g.Id)
	return
}
