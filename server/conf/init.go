package conf

import mlog "github.com/duruyao/gochat/server/log"

func init() {
	if IsNotExist() {
		if err := createFile(); err != nil {
			mlog.FatalLn(err)
		}
		if err := writeFile(&cfg); err != nil {
			mlog.FatalLn(err)
		}
	} else {
		if err := readFile(&cfg); err != nil {
			mlog.FatalLn(err)
		}
	}
}
