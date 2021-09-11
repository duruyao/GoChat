package backend

import (
	"github.com/duruyao/gochat/server/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func createGroup(c *gin.Context) {
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

func readGroup(c *gin.Context) {
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
	c.JSON(http.StatusOK, g)
}

func updateGroup(c *gin.Context) {
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

func queryGroup(c *gin.Context) {
	if c.Query("limit") != "" {
		queryGroups(c)
		return
	}
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
		//c.JSON(http.StatusBadRequest, gin.H{"error": "not found valid params in url"})
		queryGroups(c)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, g)
}

func queryGroups(c *gin.Context) {
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

func queryGroupMembers(c *gin.Context) {
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

func queryGroupAdministrator(c *gin.Context) {
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
