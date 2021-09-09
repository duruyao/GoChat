package web

import (
	mlog "github.com/duruyao/gochat/server/log"
	"github.com/duruyao/gochat/server/util"
	"os"
)

func init() {
	info1, err1 := os.Stat(TLSCertPath())
	info2, err2 := os.Stat(TLSKeyPath())
	if os.IsNotExist(err1) || os.IsNotExist(err2) ||
		(info1 != nil && info1.Size() == 0) || (info2 != nil && info2.Size() == 0) {
		mlog.ErrorLn("Not found valid TLS cert and TLS key, `util.CreateTLSCertAndKey()` will be called")
		mlog.ErrorF("Put your TLS cert here: %s\n", TLSCertPath())
		mlog.ErrorF("Put your TLS key  here: %s\n", TLSKeyPath())
		if err := util.CreateTLSCertAndKey(TLSCertPath(), TLSKeyPath()); err != nil {
			mlog.FatalLn(err)
		}
	}
}
