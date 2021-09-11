package backend

import (
	"github.com/duruyao/gochat/server/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func postUser(c *gin.Context) {
	u := data.User{}
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := u.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, u)
}

func getUsers(c *gin.Context) {
	v := c.DefaultQuery("limit", "2147483647")
	limit, err := strconv.Atoi(v)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit: " + v})
		return
	}
	us, err := data.Users(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, us)
}

func putUser(c *gin.Context) {
	u := data.User{}
	if err := c.BindUri(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := data.UserByUniqueKey("ID", u.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := u.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

func deleteUser(c *gin.Context) {
	u := data.User{}
	if err := c.BindUri(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := data.UserByUniqueKey("ID", u.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := u.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

func getUser(c *gin.Context) {
	u := data.User{}
	err := c.BindQuery(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if u.Id != 0 {
		u, err = data.UserByUniqueKey("ID", u.Id)
	} else if u.UUId != "" {
		u, err = data.UserByUniqueKey("UUID", u.UUId)
	} else if u.Name != "" {
		u, err = data.UserByUniqueKey("NAME", u.Name)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found valid params in url"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

func getUserSession(c *gin.Context) {
	u := data.User{}
	err := c.BindQuery(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if u.Id != 0 {
		u, err = data.UserByUniqueKey("ID", u.Id)
	} else if u.UUId != "" {
		u, err = data.UserByUniqueKey("UUID", u.UUId)
	} else if u.Name != "" {
		u, err = data.UserByUniqueKey("NAME", u.Name)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found valid params in url"})
		return
	}
	s, err := u.Session()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, s)
}

func getUserJoinedGroups(c *gin.Context) {
	u := data.User{}
	err := c.BindQuery(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if u.Id != 0 {
		u, err = data.UserByUniqueKey("ID", u.Id)
	} else if u.UUId != "" {
		u, err = data.UserByUniqueKey("UUID", u.UUId)
	} else if u.Name != "" {
		u, err = data.UserByUniqueKey("NAME", u.Name)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found valid params in url"})
		return
	}
	gs, err := u.JoinedGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gs)
}

func getUserCreatedGroups(c *gin.Context) {
	u := data.User{}
	err := c.BindQuery(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if u.Id != 0 {
		u, err = data.UserByUniqueKey("ID", u.Id)
	} else if u.UUId != "" {
		u, err = data.UserByUniqueKey("UUID", u.UUId)
	} else if u.Name != "" {
		u, err = data.UserByUniqueKey("NAME", u.Name)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found valid params in url"})
		return
	}
	gs, err := u.CreatedGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gs)
}

func postSession(c *gin.Context) {
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

func getSessions(c *gin.Context) {
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

func getSession(c *gin.Context) {
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
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, s)
}

func getSessionUser(c *gin.Context) {
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

func postGroup(c *gin.Context) {
	g := data.Group{}
	if err := c.BindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := data.User{Id: g.AdminId}
	g, err := u.CreateGroup(g.Name, g.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, g)
}

func getGroups(c *gin.Context) {
	v := c.DefaultQuery("limit", "2147483647")
	limit, err := strconv.Atoi(v)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit: " + v})
		return
	}
	gs, err := data.Groups(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gs)
}

func putGroup(c *gin.Context) {
	g := data.Group{}
	if err := c.BindUri(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g, err := data.GroupByUniqueKey("ID", g.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.BindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := g.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, g)
}

func deleteGroup(c *gin.Context) {
	g := data.Group{}
	if err := c.BindUri(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g, err := data.GroupByUniqueKey("ID", g.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := g.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, g)
}

func getGroup(c *gin.Context) {
	g := data.Group{}
	err := c.BindQuery(&g)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if g.Id != 0 {
		g, err = data.GroupByUniqueKey("ID", g.Id)
	} else if g.UUId != "" {
		g, err = data.GroupByUniqueKey("UUID", g.UUId)
	} else if g.Name != "" {
		g, err = data.GroupByUniqueKey("NAME", g.Name)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found valid params in url"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, g)
}

func getGroupMembers(c *gin.Context) {
	g := data.Group{}
	err := c.BindQuery(&g)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if g.Id != 0 {
		g, err = data.GroupByUniqueKey("ID", g.Id)
	} else if g.UUId != "" {
		g, err = data.GroupByUniqueKey("UUID", g.UUId)
	} else if g.Name != "" {
		g, err = data.GroupByUniqueKey("NAME", g.Name)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found valid params in url"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	us, err := g.Members()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, us)
}

func getGroupAdministrator(c *gin.Context) {
	g := data.Group{}
	err := c.BindQuery(&g)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if g.Id != 0 {
		g, err = data.GroupByUniqueKey("ID", g.Id)
	} else if g.UUId != "" {
		g, err = data.GroupByUniqueKey("UUID", g.UUId)
	} else if g.Name != "" {
		g, err = data.GroupByUniqueKey("NAME", g.Name)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found valid params in url"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	u, err := g.Administrator()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

func postMember(c *gin.Context) {
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
