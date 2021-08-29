package server

import (
	mlog "github.com/duruyao/gochat/server/log"
	"sync"
)

type beforeQuit struct {
	doOnce sync.Once
	mu     sync.Mutex
	funcs  []func() error
}

//
func (bq *beforeQuit) Append(f func() error) {
	bq.mu.Lock()
	defer bq.mu.Unlock()
	bq.funcs = append(bq.funcs, f)
}

//
func (bq *beforeQuit) Do() {
	bq.doOnce.Do(func() {
		for _, f := range bq.funcs {
			if err := f(); err != nil {
				mlog.ErrorLn(err)
			}
		}
	})
}

var BeforeQuit beforeQuit

var quit = make(chan struct{})

//
func Quit() <-chan struct{} { return quit }

//
func IsQuit() bool {
	select {
	case <-quit:
		return true
	default:
		return false
	}
}

var wantQuitOnce sync.Once

//
func WantQuit() {
	wantQuitOnce.Do(func() { close(quit) })
}
