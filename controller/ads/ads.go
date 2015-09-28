package ads
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-gonic/contrib/sessions"
//	"fmt"
	"html/template"
)

var (
	CachedCreative string
)

func init() {
	CachedCreative = ""
}

/**
	below are router functions
 */
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "dm/index.html", gin.H{})
}

func RenderCreative(c *gin.Context) {
	session := sessions.Default(c)
	c.HTML(http.StatusOK, "ad/creative_render.html", gin.H{"creative_content": template.HTML(session.Get("creative").(string)) })
}

func SaveCreative(c *gin.Context) {
	var json map[string]string
	if c.BindJSON(&json) == nil {
		session := sessions.Default(c)
		session.Set("creative", json["creative"])
		session.Save()
		CachedCreative = json["creative"]
	}

	c.JSON(http.StatusOK, "success")
}

func GetCreative(c *gin.Context) {
	session := sessions.Default(c)
	c.JSON(http.StatusOK, gin.H{"creative": session.Get("creative")})
}
