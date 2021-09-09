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
		mlog.ErrorF("Not found valid TLS Cert: %s and Key: %s\n", TLSCertPath(), TLSKeyPath())
		mlog.Error("github.com/duruyao/gochat/server/util.CreateTLSCertAndKey() will be called")
		if err := util.CreateTLSCertAndKey(TLSCertPath(), TLSKeyPath()); err != nil {
			mlog.FatalLn(err)
		}
	}
}
