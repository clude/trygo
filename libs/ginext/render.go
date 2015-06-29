package ginext

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type (
	// form binding
	HTMLNi struct {
		tplPath  string
	}

	HTML struct {
		tplPath  string
		Name     string
		Data     interface{}
	}
)

var htmlContentType = []string{"text/html; charset=utf-8"}

func LoadTemplates(engine *gin.Engine, path string) {
	err := BuildTemplate(path)
	if(err != nil){
		fmt.Println(err)
	}

	engine.HTMLRender = HTMLNi{
		tplPath: path,
	}
}

func (r HTMLNi) Instance(name string, data interface{}) render.Render {
	return HTML{
		tplPath: r.tplPath,
		Name: name,
		Data:     data,
	}
}

func (r HTML) Render(w http.ResponseWriter) error {
	w.Header()["Content-Type"] = htmlContentType
	if gin.IsDebugging() {
		BuildTemplate(r.tplPath)
	}
	return BeeTemplates[r.Name].ExecuteTemplate(w, r.Name, r.Data)
}

//func writeHeader(w http.ResponseWriter, code int, contentType string) {
//	w.Header().Set("Content-Type", contentType+"; charset=utf-8")
//	w.WriteHeader(code)
//}

//func (html HTML) Render(w http.RensponseWriter, code int, data ...interface{}) error {
//	writeHeader(w, code, "text/html")
//	file := data[0].(string)
//	obj := data[1]
//	if gin.IsDebugging() {
//		BuildTemplate(html.tplPath)
//	}
//	return BeeTemplates[file].ExecuteTemplate(w, file, obj)
//}