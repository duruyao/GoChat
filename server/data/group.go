package data

import (
	"fmt"
	"time"
)

type Group struct {
	Id        int       `db:"ID" json:"id" form:"id" uri:"id"`
	UUId      string    `db:"UUID" json:"uuid" form:"uuid"`
	Name      string    `db:"NAME" json:"name" form:"name"`
	AdminId   int       `db:"ADMIN_ID" json:"admin_id"`
	Token     string    `db:"TOKEN" json:"token"`
	CreatedAt time.Time `db:"CREATED_AT" json:"created_at"`
}

func GroupByUniqueKey(key string, value interface{}) (g Group, err error) {
	q := fmt.Sprintf("SELECT * FROM GROUPS WHERE %s = $2", key)
	err = db.Get(&g, q, value)
	return
}

func Groups(limit int) (gs []Group, err error) {
	q := `SELECT * FROM GROUPS LIMIT $1;`
	err = db.Select(&gs, q, limit)
	return
}

func (g *Group) Update() (err error) {
	q := `UPDATE GROUPS SET NAME = $1, TOKEN = $2 WHERE ID = $3;`
	_, err = db.Exec(q, g.Name, g.Token, g.Id)
	return
}

func (g *Group) Delete() (err error) {
	q1 := `DELETE FROM MEMBERS WHERE GROUP_ID = $1;`
	q2 := `DELETE FROM GROUPS WHERE ID = $1;`
	tx, err := db.Beginx()
	if err != nil {
		return
	}
	defer func() { _ = tx.Commit() }()
	if _, err = tx.Exec(q1, g.Id); err != nil {
		return
	}
	_, err = tx.Exec(q2, g.Id)
	return
}

func (g *Group) Administrator() (u User, err error) {
	q := `SELECT * FROM USERS WHERE ID = $1;`
	err = db.Get(&u, q, g.AdminId)
	return
}

func (g *Group) Members() (us []User, err error) {
	q := `SELECT U.* FROM USERS U, MEMBERS J WHERE U.ID = J.USER_ID AND J.GROUP_ID = $1`
	err = db.Select(&us, q, g.AdminId)
	return
}
