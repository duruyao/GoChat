package data

import (
	mlog "github.com/duruyao/gochat/server/log"
	"github.com/duruyao/gochat/server/util"
	"time"
)

func init() {
	if IsNotExist() {
		if err := createDb(); err != nil {
			mlog.FatalLn(err)
		}
	} else {
		if err := openDb(); err != nil {
			mlog.FatalLn(err)
		}
	}
	go goWaitCloseDb()
	//go goAutoDeleteExpiredGuest()
}

func goWaitCloseDb() {
	select {
	case <-util.Quit():
		if err := closeDb(); err != nil {
			mlog.FatalLn(err)
		}
	}
}

func goAutoDeleteExpiredGuest() {
	for {
		select {
		case <-util.Quit():
			return
		case <-time.After(10 * time.Minute):
			cmd := `BEGIN; 
DELETE FROM JOIN_ROOM WHERE USER_ID IN (SELECT ID FROM USERS WHERE MAX_ROLE > 3 AND CREATED_AT < $1); 
DELETE FROM USERS WHERE MAX_ROLE > $1 AND CREATED_AT < $2; COMMIT;`
			stmt, err := db.Prepare(cmd)
			if err == nil {
				if _, err := stmt.Exec(Admin, time.Now().Local().Add(-10*time.Hour)); err != nil {
					mlog.ErrorLn(err)
				}
				if err := stmt.Close(); err != nil {
					mlog.ErrorLn(err)
				}
			}
		}
	}
}
