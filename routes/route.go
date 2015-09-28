package routes

import (
	"cludezhu/trygo/controller"
	"cludezhu/trygo/controller/dm"
	"cludezhu/trygo/controller/ads"
	"github.com/gin-gonic/gin"
	"net/http"
)


func SetRouters(r *gin.Engine){
//	r.Static("/static", "./public")

	r.GET("/", controller.Index)
//	r.GET("/ads", controller.Ads)

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


	rst := gin.H{"title": "Title", "year": "2015", "posters": gin.H{"thumbnail": "http://i.imgur.com/UePbdph.jpg"}}
	slices := []gin.H {rst, rst, rst}

	r.GET("/react/images", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"movies":slices})
	})

	r.GET("/de", dm.Index)
	r.GET("/dsp", dm.Index)
	r.GET("/api/de/record", dm.GetRecord)
//	r.GET("/api/de/servers", dm.GetServerStatus)
//	r.GET("/api/de/updater", dm.GetUpdater)

	r.GET("/ads", ads.Index)
	r.GET("/ads/creative", ads.RenderCreative)
	r.POST("/api/creative/save", ads.SaveCreative)
	r.GET("/api/creative/get", ads.GetCreative)

}

