package backend

import (
	mlog "github.com/duruyao/gochat/server/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GoRunWebServer(addr string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	v1 := router.Group("v1")

	{
		v1.POST("/user", createtUser)                          // POST   /user
		v1.GET("/user/:id", readUser)                          // GET    /user/:id
		v1.PUT("/user/:id", updateUser)                        // PUT    /user/:id
		v1.DELETE("/user/:id", deleteUser)                     // DELETE /user/:id
		v1.GET("/user", queryUser)                             // GET    /user?limit=&id=&uuid=&name=
		v1.GET("/user_session", queryUserSession)              // GET    /user_session?id=&uuid=&name=
		v1.GET("/user_joined_groups", queryUserJoinedGroups)   // GET    /user_joined_groups?id=&uuid=&name=
		v1.GET("/user_created_groups", queryUserCreatedGroups) // GET    /user_created_groups?id=&uuid=&name=

		v1.POST("/session", createSession)         // POST   /session
		v1.GET("/session/:id", readSession)        // GET    /session/:id
		v1.DELETE("/session/:id", deleteSession)   // DELETE /session/:id
		v1.GET("/session", querySession)           // GET    /session?limit=&id=&uuid=&user_id=
		v1.GET("/session_user", qurerySessionUser) // GET    /session_user?id=&uuid=&user_id=

		v1.POST("/group", createGroup)                          // POST   /group
		v1.GET("/group/:id", readGroup)                         // GET    /group/:id
		v1.PUT("/group/:id", updateGroup)                       // PUT    /group/:id
		v1.DELETE("/group/:id", deleteGroup)                    // DELETE /group/:id
		v1.GET("/group", queryGroup)                            // GET    /group?limit=&id=&uuid=&name=
		v1.GET("/group_members", queryGroupMembers)             // GET    /group_members?id=&uuid=&name=
		v1.GET("/group_administrator", queryGroupAdministrator) // GET    /group_administrator?id=&uuid=&name=

		v1.POST("/member", createMember)                      // POST   /member
		v1.DELETE("/member/:group_id/:user_id", deleteMember) // DELETE /member/:group_id/:user_id
	}

	backendServer := http.Server{
		Addr:              addr,
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

	mlog.FatalLn(backendServer.ListenAndServe())
}
