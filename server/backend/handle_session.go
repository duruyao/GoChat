package backend

import (
	"github.com/duruyao/gochat/server/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func createSession(c *gin.Context) {
	s := data.Session{}
	if err := c.BindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := data.User{Id: s.UserId}
	s, err := u.CreateSession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, s)
}

func readSession(c *gin.Context) {
	s := data.Session{}
	if err := c.BindUri(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s, err := data.SessionByUniqueKey("ID", s.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, s)
}

func deleteSession(c *gin.Context) {
	s := data.Session{}
	if err := c.BindUri(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s, err := data.SessionByUniqueKey("ID", s.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, s)
}

func querySessions(c *gin.Context) {
	v := c.DefaultQuery("limit", "2147483647")
	limit, err := strconv.Atoi(v)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit: " + v})
		return
	}
	ss, err := data.Sessions(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ss)
}

func querySession(c *gin.Context) {
	if c.Query("limit") != "" {
		querySessions(c)
		return
	}
	s := data.Session{}
	err := c.BindQuery(&s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if s.Id != 0 {
		s, err = data.SessionByUniqueKey("ID", s.Id)
	} else if s.UUId != "" {
		s, err = data.SessionByUniqueKey("UUID", s.UUId)
	} else if s.UserId != 0 {
		s, err = data.SessionByUniqueKey("USER_ID", s.UserId)
	} else {
		//c.JSON(http.StatusBadRequest, gin.H{"error": "not found valid params in url"})
		querySessions(c)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, s)
}

func qurerySessionUser(c *gin.Context) {
	s := data.Session{}
	err := c.BindQuery(&s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if s.Id != 0 {
		s, err = data.SessionByUniqueKey("ID", s.Id)
	} else if s.UUId != "" {
		s, err = data.SessionByUniqueKey("UUID", s.UUId)
	} else if s.UserId != 0 {
		s, err = data.SessionByUniqueKey("USER_ID", s.UserId)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found valid params in url"})
		return
	}
	u, err := s.User()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}
