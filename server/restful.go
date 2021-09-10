package server

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

func getUser(c *gin.Context) {
	k, v := "", ""
	if v = c.Query("id"); len(v) > 0 {
		k = "ID"
		if _, err := strconv.Atoi(v); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id: " + v})
			return
		}
	} else if v = c.Query("uuid"); len(v) > 0 {
		k = "UUID"
	} else if v = c.Query("name"); len(v) > 0 {
		k = "NAME"
	}
	u, err := data.UserByUniqueKey(k, v)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id: " + c.Param("id")})
		return
	}
	u, err := data.UserByUniqueKey("ID", id)
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id: " + c.Param("id")})
		return
	}
	u, err := data.UserByUniqueKey("ID", id)
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
