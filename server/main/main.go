package main

import (
	"github.com/duruyao/gochat/server/backend"
	"github.com/duruyao/gochat/server/util"
	"github.com/duruyao/gochat/server/web"
)

func main() {
	go web.GoRunWebApp()
	go backend.GoRunWebServer(":1213")
	go backend.GoRunChatServer(":1314")

	select {
	case <-util.Quit():
		return
	}
}
