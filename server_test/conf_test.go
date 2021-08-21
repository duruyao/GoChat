package server_test

import (
	"github.com/duruyao/gochat/server/conf"
	"testing"
)

func TestConfCtCreateReadWrite(t *testing.T) {
	if err := conf.CreateFile(); err != nil {
		t.Fatal(err)
	}
	var cf1 conf.Conf
	if err := conf.ReadFile(&cf1); err != nil {
		t.Fatal(err)
	}
	t.Log(cf1.String())
	cf1 = conf.DefaultConf()
	if err := conf.WriteFile(&cf1); err != nil {
		t.Fatal(err)
	}
	var cf2 conf.Conf
	if err := conf.ReadFile(&cf2); err != nil {
		t.Fatal(err)
	}
	t.Log(cf2.String())
}

func TestConfReadwrite(t *testing.T) {
	var cf1 conf.Conf
	if err := conf.ReadFile(&cf1); err != nil {
		t.Fatal(err)
	}
	t.Log(cf1.String())
	cf1.AppendAdmins([]conf.User{{"admin1", "123456"}, {"admin2", "654321"}}...)
	if err := conf.WriteFile(&cf1); err != nil {
		t.Fatal(err)
	}
	var cf2 conf.Conf
	if err := conf.ReadFile(&cf2); err != nil {
		t.Fatal(err)
	}
	t.Log(cf2.String())
}
