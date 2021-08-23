package server

import (
	sev "github.com/duruyao/gochat/server"
	"testing"
)

func TestInit(t *testing.T) {
	defer func() {
		if err := sev.Quit(); err != nil {
			sev.Fatal(err)
		}
	}()
	sev.Info("Load configuration:\n" + sev.Config.String())
	sev.Info("Load database:\n" + sev.Table.String())
}
