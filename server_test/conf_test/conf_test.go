package conf_test

import (
	"github.com/duruyao/gochat/server/conf"
	"testing"
)

func TestConfCtCreateReadWrite(t *testing.T) {
	if err := conf.CreateFile(); err != nil {
		t.Fatal(err)
	}
	config := conf.NewConfig()
	if err := conf.ReadFile(config); err != nil {
		t.Fatal(err)
	}
	t.Log(config.String())
	config = conf.NewDefaultConfig()
	if err := conf.WriteFile(config); err != nil {
		t.Fatal(err)
	}
	if err := conf.ReadFile(config); err != nil {
		t.Fatal(err)
	}
	t.Log(config.String())
}

func TestConfReadwrite(t *testing.T) {
	config := conf.NewConfig()
	if err := conf.ReadFile(config); err != nil {
		t.Fatal(err)
	}
	t.Log(config.String())
	config.SetMaxUsers(30000)
	config.SetRoot("12345678")
	if err := conf.WriteFile(config); err != nil {
		t.Fatal(err)
	}
	if err := conf.ReadFile(config); err != nil {
		t.Fatal(err)
	}
	t.Log(config.String())
}
