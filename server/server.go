package server

import (
	mlog "github.com/duruyao/gochat/server/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/user", postUser)
	router.GET("/user", getUser)
	router.GET("/users", getUsers)
	router.PUT("/user/:id", putUser)
	router.DELETE("/user/:id", deleteUser)

	myServer := http.Server{
		Addr:              ":1213",
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       7 * time.Second,
		ReadHeaderTimeout: 7 * time.Second,
		WriteTimeout:      7 * time.Second,
		IdleTimeout:       7 * time.Second,
		MaxHeaderBytes:    1024 * 1024,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          mlog.ErrorLogger,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	mlog.FatalLn(myServer.ListenAndServe())
}
