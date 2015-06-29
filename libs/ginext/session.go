package ginext

import(
//	"encoding/gob"
//	"github.com/gin-gonic/gin"
)

type Person struct {
	FirstName    string
	LastName     string
	Email        string
	Age            int
}

func init() {
//	gob.Register(&Person{})
//	gob.Register(&gin.H{})
}