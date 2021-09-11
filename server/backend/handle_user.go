package backend

import (
	"github.com/duruyao/gochat/server/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func createtUser(c *gin.Context) {
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

func readUser(c *gin.Context) {
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
	c.JSON(http.StatusOK, u)
}

func updateUser(c *gin.Context) {
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

func queryUsers(c *gin.Context) {
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

func queryUser(c *gin.Context) {
	if c.Query("limit") != "" {
		queryUsers(c)
		return
	}
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
		//c.JSON(http.StatusBadRequest, gin.H{"error": "not found valid params in url"})
		queryUsers(c)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

func queryUserSession(c *gin.Context) {
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

func queryUserJoinedGroups(c *gin.Context) {
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

func queryUserCreatedGroups(c *gin.Context) {
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
