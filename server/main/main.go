package main

import (
	"github.com/duruyao/gochat/server/backend"
	"github.com/duruyao/gochat/server/util"
	"github.com/duruyao/gochat/server/web"
)

func main() {
	go web.GoRunWebApp()
	go backend.GoRunWebServer(":8181")
	go backend.GoRunChatServer(":8282")

	select {
	case <-util.Quit():
		return
	}
}
