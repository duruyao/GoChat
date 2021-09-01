package db

import "testing"

func Test(t *testing.T) {

}

func TestCreateDB(t *testing.T) {
	if err := CreateDB(); err != nil {
		t.Fatal(err)
	}
}

func TestAdminsTable_Query(t *testing.T) {
	if err := OpenDB(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := CloseDB(); err != nil {
			t.Fatal(err)
		}
	}()
	adminsTable := new(AdminsTable)
	admins, err := adminsTable.Query()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(admins)
}

func TestAdminsTable_Insert(t *testing.T) {
	if err := OpenDB(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := CloseDB(); err != nil {
			t.Fatal(err)
		}
	}()
	adminsTable := new(AdminsTable)
	if err := adminsTable.Insert(Admin{Uid: "alice", Pwd: "dhjkahdjkl"}); err != nil {
		t.Fatal(err)
	}
	if err := adminsTable.Insert(Admin{Uid: "bob", Pwd: "ddhkducnah"}); err != nil {
		t.Fatal(err)
	}
}

func TestAdminsTable_Delete(t *testing.T) {
	if err := OpenDB(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := CloseDB(); err != nil {
			t.Fatal(err)
		}
	}()
	adminsTable := new(AdminsTable)
	if err := adminsTable.Delete("alice"); err != nil {
		t.Fatal(err)
	}
}

func TestRoomsTable_Query(t *testing.T) {
	if err := OpenDB(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := CloseDB(); err != nil {
			t.Fatal(err)
		}
	}()
	roomsTable := new(RoomsTable)
	rooms, err := roomsTable.Query()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rooms)
}

func TestRoomsTable_Insert(t *testing.T) {
	if err := OpenDB(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := CloseDB(); err != nil {
			t.Fatal(err)
		}
	}()
	roomsTable := new(RoomsTable)
	if err := roomsTable.Insert(Room{Rid: "125", Admin: "alice", Token: "ahvpzidjsy"}); err != nil {
		t.Fatal(err)
	}
	if err := roomsTable.Insert(Room{Rid: "523", Admin: "chris", Token: "alkovuytfg"}); err != nil {
		t.Fatal(err) // NOTE: chris is not in ADMINS_TB
	} // TODO: find BUG about REFERENCES in SQLite
}

func TestRoomsTable_Delete(t *testing.T) {
	if err := OpenDB(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := CloseDB(); err != nil {
			t.Fatal(err)
		}
	}()
	roomsTable := new(RoomsTable)
	if err := roomsTable.Delete("abc"); err != nil {
		t.Error(err)
	}
	if err := roomsTable.Delete("alice"); err != nil {
		t.Fatal(err)
	}
}
