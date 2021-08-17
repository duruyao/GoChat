package server_test

import (
	"fmt"
	"testing"

	chatserver "github.com/duruyao/gochat/server"
)

func Test(t *testing.T) {
	myConf := chatserver.NewConf()
	myConf.Load()
	fmt.Println(myConf)
	myConf.AddAdmin("admin1", "123456")
	myConf.AddAdmin("admin2", "654321")
	fmt.Println(myConf)
	myConf.Save()
	myConf.Load()
	fmt.Println(myConf)
}
