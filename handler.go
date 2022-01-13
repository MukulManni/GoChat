package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	username = "user"
	ucolor   = "color"
)

func loadLogin(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{},
	)
}

func auth(c *gin.Context) {

	user := c.PostForm("user")
	color := c.PostForm("Color")

	if strings.Contains(user, "<") || strings.Contains(user, ">") || user == "" {
		c.HTML(
			http.StatusOK,
			"login.html",
			gin.H{
				"error": "Use a valid username",
			},
		)
		return
	}

	session := sessions.Default(c)

	session.Set(username, user)
	session.Set(ucolor, color)

	if err := session.Save(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": "Unable to save session."})
		return

	} else {
		c.Redirect(http.StatusFound, "/u/chat")
	}
}

func isLogin(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(username)

	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		return
	}

	c.Next()
}

func chatpage(c *gin.Context) {

	session := sessions.Default(c)
	user := session.Get(username)
	color := session.Get(ucolor)
	msglist := getAllMsgs()

	c.HTML(
		http.StatusOK,
		"chat.html",
		gin.H{
			"user":    user,
			"color":   color,
			"msgList": msglist,
		},
	)

}

func postmsg(c *gin.Context) {

	session := sessions.Default(c)
	user := session.Get(username)
	color := session.Get(ucolor)
	time := time.Now()

	umessage := c.PostForm("usermessage")

	data, err := addMsgtoDB(db, fmt.Sprint(user), umessage, fmt.Sprint(color), time.Format("[15:04:05] "))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Unable to add message to database"})
	}

	addMsg(*data)

	c.Redirect(http.StatusFound, "/u/chat")
}

func jsonmsg(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		globalmsgList,
	)
}
