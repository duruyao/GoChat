package server

import (
	sev "github.com/duruyao/gochat/server"
	"testing"
)

func TestInit(t *testing.T) {
	defer sev.BeforeQuit.Do()
	defer sev.WantQuit()
	sev.Info("Load configuration:\n" + sev.Config.String())
	sev.Info("Load database:\n" + sev.Table.String())
}
