package server_test

import (
	"github.com/duruyao/gochat/server/db"
	"testing"
)

func TestDbCreate(t *testing.T) {
	if err := db.CreateFile(); err != nil {
		t.Fatal(err)
	}
}

func TestDbReadWrite(t *testing.T) {
	{
		tab := db.NewTable()
		tab.Insert(db.Data{Room: "237", Uid: "admin1", Token: "adjakdhjad"})
		tab.Insert(db.Data{Room: "238", Uid: "admin2", Token: "ahdjkagfjk"})
		t.Log(tab.String())
		if err := db.WriteFile(tab); err != nil {
			t.Fatal(err)
		}
	}
	{
		tab := db.NewTable()
		if err := db.ReadFile(tab); err != nil {
			t.Fatal(err)
		}
		t.Log(tab.String())
		tab.Insert(db.Data{Room: "239", Uid: "admin3", Token: "dhdagkffga"})
		tab.Insert(db.Data{Room: "240", Uid: "admin4", Token: "dalhfglahf"})
		if err := db.WriteFile(tab); err != nil {
			t.Fatal(err)
		}
		if err := db.ReadFile(tab); err != nil {
			t.Fatal(err)
		}
		t.Log(tab.String())
	}
}
