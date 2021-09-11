package backend

import (
	"github.com/duruyao/gochat/server/data"
	"github.com/gin-gonic/gin"
	"net/http"
)


func createMember(c *gin.Context) {
	m := struct {
		UserId  int `json:"user_id"`
		GroupId int `json:"group_id"`
	}{}
	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := data.User{Id: m.UserId}
	g := data.Group{Id: m.GroupId}
	if err := u.JoinGroup(g); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusCreated, "")
}

func deleteMember(c *gin.Context) {
	m := struct {
		UserId  int `uri:"user_id" binding:"required"`
		GroupId int `uri:"group_id" binding:"required"`
	}{}
	if err := c.BindUri(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := data.User{Id: m.UserId}
	g := data.Group{Id: m.GroupId}
	if err := u.LeaveGroup(g); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "")
}
