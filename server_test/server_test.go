package server_test

import (
	"fmt"
	"github.com/duruyao/gochat/server"
	"testing"
)

func TestConf(t *testing.T) {
	cfg := server.GetConfigurator()
	cfg.LoadConf()
	fmt.Println(cfg)
	conf := cfg.Conf()
	conf.AddAdmin(server.Admin{Uid: "admin3", Pwd: "123456"})
	conf.AddAdmin(server.Admin{Uid: "admin4", Pwd: "654321"})
	cfg.SetConf(&conf)
	fmt.Println(cfg)
	cfg.SaveConf()
	cfg.LoadConf()
	fmt.Println(cfg)
}
