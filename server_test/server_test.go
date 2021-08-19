package server_test

import (
	"fmt"
	chatserver "github.com/duruyao/gochat/server"
	"testing"
)

func TestConf(t *testing.T) {
	cfg := chatserver.GetConfigurator()
	cfg.LoadConf()
	fmt.Println(cfg)
	conf := cfg.Conf()
	conf.AddAdmin(chatserver.Admin{Uid: "admin3", Pwd: "123456"})
	conf.AddAdmin(chatserver.Admin{Uid: "admin4", Pwd: "654321"})
	cfg.SetConf(&conf)
	fmt.Println(cfg)
	cfg.SaveConf()
	cfg.LoadConf()
	fmt.Println(cfg)
}

func TestLog(t *testing.T) {
	cfg := chatserver.GetConfigurator()
	conf := cfg.Conf()
	logger := chatserver.GetLogger(conf.LogFileEnable)
	server := chatserver.GetServer(cfg, logger)
	server.Start()
	server.Stop()
	server.Pause()
	server.Resume()
	server.Restart()
	server.Pause()
	server.Resume()
	server.Stop()
}
