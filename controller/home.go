package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

)

type greetings struct {
	Intro    string
	Messages []string
}


func Index(c *gin.Context) {
	passedObj := greetings{
		Intro:    "Hello from Go!",
		Messages: []string{"Hello!", "Hi!", "¡Hola!", "Bonjour!", "Ciao!", "<script>evilScript()</script>"},
	}

	c.HTML(http.StatusOK, "ads/demo.html", passedObj)
}


func Ads(c *gin.Context) {
	c.HTML(http.StatusOK, "ads/demo.html",
		gin.H{
			"Intro": "Hello",
			"Messages": []string{"My!", "Name!", "¡Is!","WFFF"},
		})
}
