package web

import (
	mlog "github.com/duruyao/gochat/server/log"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func projectDir() string {
	dir, err := os.Getwd()
	if err != nil {
		mlog.FatalLn(err)
	}
	for {
		base := filepath.Base(dir)
		if strings.HasPrefix(base, "gochat") || strings.HasPrefix(base, "GoChat") {
			return dir
		}
		if filepath.Dir(dir) == dir {
			break
		}
		dir = filepath.Dir(dir)
	}
	mlog.FatalLn("not found project directory")
	return ""
}

func HtmlDir() string {
	return projectDir() + "/html"
}

func CssDir() string {
	return projectDir() + "/css"
}

func JsDir() string {
	return projectDir() + "/js"
}
func ImgDir() string {
	return projectDir() + "/img"
}

func AllSourcePaths() []string {
	var paths []string
	for _, dir := range []string{CssDir(), JsDir(), ImgDir()} {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			mlog.ErrorLn(err)
			return nil
		}
		reg := regexp.MustCompile(`.+\.(css|js|json|xml|txt|svg|png|jpg|jpeg)$`)
		if reg == nil {
			mlog.FatalLn("regexp error")
			return nil
		}
		for _, file := range files {
			if reg.FindString(file.Name()) != "" {
				paths = append(paths, dir+"/"+file.Name())
			}
		}
	}
	return paths
}

func AllHtmlPaths() []string {
	var paths []string
	files, err := ioutil.ReadDir(HtmlDir())
	if err != nil {
		mlog.ErrorLn(err)
		return nil
	}
	reg := regexp.MustCompile(`.+\.(html|HTML)$`)
	if reg == nil {
		mlog.FatalLn("regexp error")
		return nil
	}
	for _, file := range files {
		if reg.FindString(file.Name()) != "" {
			paths = append(paths, HtmlDir()+"/"+file.Name())
		}
	}
	return paths
}

func filenameWithoutExt(filename string) string {
	name := filepath.Base(filename)
	return name[:len(name)-len(filepath.Ext(name))]
}
