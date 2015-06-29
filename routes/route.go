package routes

import (
	"../controller"
	"github.com/gin-gonic/gin"
	"net/http"
)


func SetRouters(r *gin.Engine){
//	r.Static("/static", "./public")

	r.GET("/", controller.Index)
	r.GET("/ads", controller.Ads)

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/join/
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		action := c.Request.URL.Query().Get("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.GET("/incr", controller.TestSessionUser)
	r.GET("/incrg", controller.TestSessionGinH)
}

