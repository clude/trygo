package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"encoding/gob"
)




// Options stores configuration for a session or session store.
// Fields are a subset of http.Cookie fields.
type User struct {
	Id int
	Name   string
	Title string
	Counter int
}

func init(){
	gob.Register(User{})
	gob.Register(gin.H{})
}

func TestSessionUser(c *gin.Context) {
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count += 1
	}

	session.Set("count", count)
	user := User{Id:123, Name:"tester", Title:"Mgr", Counter: count}
	session.Set("user", user)
	session.Save()

	user2 := session.Get("user").(User);

	c.JSON(200, gin.H{"count": user2.Counter})
}

func TestSessionGinH(c *gin.Context) {
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count += 1
	}

	session.Set("count", count)
	session.Set("gh", gin.H{"Id":123, "Counter":count})
	session.Save()

	user2 := session.Get("gh").(gin.H);
	c.JSON(200, gin.H{"count": user2["Counter"]})
}