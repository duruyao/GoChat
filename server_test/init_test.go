package server

import (
	sev "github.com/duruyao/gochat/server"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	defer sev.BeforeQuit.Do()
	defer sev.WantQuit()
	sev.Info("Load configuration:\n" + sev.MyConfig.String())
	sev.Info("Load database:\n" + sev.MyRoomTable.String())
	time.Sleep(3 * time.Second)
}
