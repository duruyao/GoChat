package db_test

import (
	"github.com/duruyao/gochat/server/db"
	"testing"
)

func TestDbCreateQuery(t *testing.T) {
	if err := db.CreateDB(); err != nil {
		t.Fatal(err)
	}
	if err := db.OpenDB(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := db.CloseDB(); err != nil {
			t.Fatal(err)
		}
	}()
	adminsTable := new(db.AdminsTable)
	admins, err := adminsTable.Query()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(admins)
	roomsTable := new(db.RoomsTable)
	rooms, err := roomsTable.Query()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rooms)
}

func TestDbQuery(t *testing.T) {
	if err := db.OpenDB(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := db.CloseDB(); err != nil {
			t.Fatal(err)
		}
	}()
	adminsTable := new(db.AdminsTable)
	admins, err := adminsTable.Query()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(admins)
	roomsTable := new(db.RoomsTable)
	rooms, err := roomsTable.Query()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rooms)
}

func TestDbInsertDeleteQuery(t *testing.T) {
	if err := db.CreateDB(); err != nil {
		t.Fatal(err)
	}
	if err := db.OpenDB(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := db.CloseDB(); err != nil {
			t.Fatal(err)
		}
	}()
	adminsTable := new(db.AdminsTable)
	admins, err := adminsTable.Query()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(admins)
	roomsTable := new(db.RoomsTable)
	rooms, err := roomsTable.Query()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rooms)

	if err := adminsTable.Insert(db.Admin{Uid: "alice", Pwd: "dhjkahdjkl"}); err != nil {
		t.Fatal(err)
	}
	if err := adminsTable.Insert(db.Admin{Uid: "bob", Pwd: "ddhkducnah"}); err != nil {
		t.Fatal(err)
	}
	admins, err = adminsTable.Query()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(admins)

	if err := roomsTable.Insert(db.Room{Rid: "125", Admin: "alice", Token: "ahvpzidjsy"}); err != nil {
		t.Fatal(err)
	}
	if err := roomsTable.Insert(db.Room{Rid: "523", Admin: "chris", Token: "alkovuytfg"}); err != nil {
		t.Error(err) // NOTE: chris is not in ADMINS_TB
	} // TODO: find BUG about REFERENCES in SQLite
	rooms, err = roomsTable.Query()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rooms)
}
