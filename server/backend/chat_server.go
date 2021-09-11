package backend

import (
	mlog "github.com/duruyao/gochat/server/log"
	mmsg "github.com/duruyao/gochat/server/msg"
	"net"
)

type ClientCh chan<- mmsg.Message

var (
	joiningCh  = make(chan ClientCh)
	leavingCh  = make(chan ClientCh)
	messageCh  = make(chan mmsg.Message)
	clientsChs = make(map[ClientCh]bool)
)

func GoRunChatServer(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		mlog.FatalLn(err)
	}

	go goBroadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			mlog.ErrorLn(err)
			continue
		}
		go goHandleConn(conn)
	}
}
