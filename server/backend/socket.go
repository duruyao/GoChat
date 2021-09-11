package backend

import (
	"github.com/duruyao/gochat/server/data"
	mlog "github.com/duruyao/gochat/server/log"
	mmsg "github.com/duruyao/gochat/server/msg"
	"net"
	"time"
)

func goBroadcast() {
	for {
		select {
		case msg := <-messageCh:
			for cli := range clientsChs {
				cli <- msg
			}
		case cli := <-joiningCh:
			clientsChs[cli] = true
		case cli := <-leavingCh:
			delete(clientsChs, cli)
			close(cli)
		}
	}
}

func goHandleConn(conn net.Conn) {
	var client mmsg.Sender

	ch := make(chan mmsg.Message)

	defer func() {
		leavingCh <- ch
		if err := conn.Close(); err != nil {
			mlog.ErrorLn(err)
		}
	}()

	go func() {
		for msg := range ch {
			if groupHasMemberByName(msg.GroupName(), client.Name) {
				msg.SetForwardTime(time.Now().Local())
				bytes, _ := msg.Serialize()
				if _, err := conn.Write(bytes); err != nil {
					mlog.ErrorLn(err)
					return
				}
			}
		}
	}()

	joiningCh <- ch

	for {
		var bytes []byte
		msg := mmsg.Message{}
		if _, err := conn.Read(bytes); err != nil {
			mlog.ErrorLn(err)
			return
		}
		if msg.Parse(bytes) != nil {
			client = msg.Sender()
			messageCh <- msg
		}
	}
}

func groupHasMemberByName(groupName string, userName string) bool {
	g, _ := data.GroupByUniqueKey("NAME", groupName)
	u, _ := data.UserByUniqueKey("NAME", userName)
	return g.HasMember(u)
}
