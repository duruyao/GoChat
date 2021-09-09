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
	//go goAutoDeleteExpiredGuest(30 * time.Hour)
}

func goWaitCloseDb() {
	select {
	case <-util.Quit():
		if err := closeDb(); err != nil {
			mlog.FatalLn(err)
		}
	}
}

func goAutoDeleteExpiredGuest(effectiveTime time.Duration) {
	for {
		select {
		case <-util.Quit():
			return
		case <-time.After(10 * time.Minute):
			q1 := `DELETE FROM MEMBERS WHERE USER_ID IN (SELECT ID FROM USERS WHERE MAX_ROLE > 3 AND CREATED_AT < $1);`
			q2 := `DELETE FROM SESSIONS WHERE USER_ID IN (SELECT ID FROM USERS WHERE MAX_ROLE > 3 AND CREATED_AT < $1);`
			q3 := `DELETE FROM USERS WHERE MAX_ROLE > 3 AND CREATED_AT < $1;`
			tx, err := db.Begin()
			if err == nil {
				date := time.Now().Local().Add(-effectiveTime)
				if _, err := tx.Exec(q1, date); err != nil {
					mlog.ErrorLn(err)
					continue
				}
				if _, err := tx.Exec(q2, date); err != nil {
					mlog.ErrorLn(err)
					continue
				}
				if _, err := tx.Exec(q3, date); err != nil {
					mlog.ErrorLn(err)
					continue
				}
				if err := tx.Commit(); err != nil {
					mlog.ErrorLn(err)
				}
			}
		}
	}
}
