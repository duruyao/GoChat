package db

import (
	"database/sql"
)

var db *sql.DB

type RoomsTable struct{}

type Room struct {
	Rid   string `json:"room_id"`
	Admin string `json:"admin_id"`
	Token string `json:"token"`
}

//
func (t *RoomsTable) Query() ([]Room, error) {
	var rooms []Room
	rows, err := db.Query(roomsQuerySQL)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var room Room
		if err := rows.Scan(&room.Rid, &room.Admin, &room.Token); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rooms, nil
}

//
func (t *RoomsTable) Insert(room Room) (err error) {
	var stmt *sql.Stmt
	//stmt, e = db.Prepare(roomsCreateSQL)
	//if e != nil {
	//	return e // TODO: note the bug: return err(nil) instead of e
	//}
	stmt, err = db.Prepare(roomsInsertSQL) // avoid SQL injections
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(room.Rid, room.Admin, room.Token); err != nil {
		return err
	}
	defer func() { _ = stmt.Close() }()
	return err
}

//
func (t *RoomsTable) Delete(rid string) (err error) {
	var stmt *sql.Stmt
	stmt, err = db.Prepare(roomsDeleteSQL) // avoid SQL injections
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(rid); err != nil {
		return err
	}
	defer func() { _ = stmt.Close() }()
	return err
}

//
func (t *RoomsTable) ExecOneSQL(query string, args ...interface{}) (err error) {
	var stmt *sql.Stmt
	stmt, err = db.Prepare(query) // avoid SQL injections
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(args...); err != nil {
		return err
	}
	defer func() { _ = stmt.Close() }()
	return err
}

type User struct {
	Uid string `json:"uid"`
	Pwd string `json:"pwd"`
}

type Admin User

type AdminsTable struct{}

//
func (t *AdminsTable) Query() ([]Admin, error) {
	var admins []Admin
	rows, err := db.Query(adminsQuerySQL)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var admin Admin
		if err := rows.Scan(&admin.Uid, &admin.Pwd); err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return admins, nil
}

//
func (t *AdminsTable) Insert(admin Admin) (err error) {
	var stmt *sql.Stmt
	stmt, err = db.Prepare(adminsInsertSQL) // avoid SQL injections
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(admin.Uid, admin.Pwd); err != nil {
		return err
	}
	defer func() { _ = stmt.Close() }()
	return err
}

//
func (t *AdminsTable) Delete(uid string) (err error) {
	var stmt *sql.Stmt
	stmt, err = db.Prepare(adminsDeleteSQL) // avoid SQL injections
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(uid); err != nil {
		return err
	}
	defer func() { _ = stmt.Close() }()
	return err
}

//
func (t *AdminsTable) ExecOneSQL(query string, args ...interface{}) (err error) {
	var stmt *sql.Stmt
	stmt, err = db.Prepare(query) // avoid SQL injections
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(args...); err != nil {
		return err
	}
	defer func() { _ = stmt.Close() }()
	return err
}

//
func createRoomsTable() error {
	if _, err := db.Exec(roomsCreateSQL); err != nil {
		return err
	}
	return nil
}

//
func createAdminsTable() error {
	if _, err := db.Exec(adminsCreateSQL); err != nil {
		return err
	}
	return nil
}
