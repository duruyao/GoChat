package web

import (
	"testing"
)

func TestAllHtmlPaths(t *testing.T) {
	t.Log(AllHtmlPaths())
}

func TestGoHttpServer(t *testing.T) {
	GoHttpService(":1213")
}

func TestGoHttpsService(t *testing.T) {
	GoHttpsService(":1213")
}
