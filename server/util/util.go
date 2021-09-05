package util

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Tomorrow() returns an object of type time.Time whose value is 00:00:00.000 of the next day.
func Tomorrow() time.Time {
	t := time.Now().Local().AddDate(0, 0, 1)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

var gochatDirOnce sync.Once
var gochatDir string

// GoChatDir returns the directory of the current project, such as "/home/user/project/gochat*".
func GoChatDir() string {
	gochatDirOnce.Do(func() {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		for {
			base := filepath.Base(dir)
			if strings.HasPrefix(base, "gochat") || strings.HasPrefix(base, "GoChat") {
				gochatDir = dir
				return
			}
			if filepath.Dir(dir) == dir {
				break
			}
			dir = filepath.Dir(dir)
		}
		log.Fatalln("not found project directory")
	})
	return gochatDir
}

var userHomeDirOnce sync.Once
var userHomeDir string

// GoChatDir returns the home directory of the current user, such as "/home/user" in Unix-like OS.
func UserHomeDir() string {
	userHomeDirOnce.Do(func() {
		var err error
		userHomeDir, err = os.UserHomeDir()
		if err != nil {
			log.Fatalln(err)
		}
	})
	return userHomeDir
}

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

var quitOnce sync.Once

//
func SetQuit() { quitOnce.Do(func() { close(quit) }) }

// CreateUUId creates a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func CreateUUId() string {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

func Encrypt(plaintext string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(plaintext)))
}

func WithoutExt(filename string) string {
	name := filepath.Base(filename)
	return name[:len(name)-len(filepath.Ext(name))]
}
