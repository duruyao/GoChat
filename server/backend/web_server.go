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
		v1.POST("/user", postUser)                           // POST   /user
		v1.GET("/users", getUsers)                           // GET    /users?limit=...
		v1.PUT("/user/:id", putUser)                         // PUT    /user/...
		v1.DELETE("/user/:id", deleteUser)                   // DELETE /user/...
		v1.GET("/user", getUser)                             // GET    /user?id=...&uuid=...&name=...
		v1.GET("/user_session", getUserSession)              // GET    /user_session?id=...&uuid=...&name=...
		v1.GET("/user_joined_groups", getUserJoinedGroups)   // GET    /user_joined_groups?id=...&uuid=...&name=...
		v1.GET("/user_created_groups", getUserCreatedGroups) // GET    /user_created_groups?id=...&uuid=...&name=...

		v1.POST("/session", postSession)         // POST   /session
		v1.GET("/sessions", getSessions)         // GET    /sessions?limit=...
		v1.DELETE("/session/:id", deleteSession) // DELETE /session/...
		v1.GET("/session", getSession)           // GET    /session?id=...&uuid=...&user_id=...
		v1.GET("/session_user", getSessionUser)  // GET    /session_user?id=...&uuid=...&user_id=...

		v1.POST("/group", postGroup)                          // POST   /group
		v1.GET("/groups", getGroups)                          // GET    /groups?limit=...
		v1.PUT("/group/:id", putGroup)                        // PUT    /group/...
		v1.DELETE("/group/:id", deleteGroup)                  // DELETE /group/...
		v1.GET("/group", getGroup)                            // GET    /group?id=...&uuid=...&name=...
		v1.GET("/group_members", getGroupMembers)             // GET    /group_members?id=...&uuid=...&name=...
		v1.GET("/group_administrator", getGroupAdministrator) // GET    /group_administrator?id=...&uuid=...&name=...

		v1.POST("/member", postMember)                      // POST   /member
		v1.DELETE("/member/:group_id/:user_id", postMember) // DELETE /member/.../...
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
