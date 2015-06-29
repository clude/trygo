package main

import (
	"./routes"
	"./libs/ginext"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
//	"log"
//	"net/http"
	"os"
	"path/filepath"
//	"time"
//	"github.com/unrolled/render"
	"github.com/gin-gonic/contrib/sessions"

)

var templateDelims = []string{"{{%", "%}}"}

var templates *template.Template

//var rd *render.Render


func initBak() {
	// initialize the templates,
	// couldn't have used http://golang.org/pkg/html/template/#ParseGlob
	// since we have custom delimiters.
	basePath := "tpl/"
	fmt.Println(basePath)
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// don't process folders themselves
		if info.IsDir() {
			return nil
		}
		templateName := path[len(basePath):]
		if templates == nil {
			templates = template.New(templateName)
			templates.Delims(templateDelims[0], templateDelims[1])
			_, err = templates.ParseFiles(path)
		} else {
			_, err = templates.New(templateName).ParseFiles(path)
		}


		fmt.Printf("Processed template %s %s\n", path, templateName)

		for k, v := range templates.Templates() {
			fmt.Printf("Processed template %s %s\n", k, v.Name())
		}
		return err
	})
	if err != nil {
		fmt.Print(err)
	}
}

func main() {
	fmt.Println("Booting up the server....")
//	gin.SetMode("release")
	r := gin.Default()

	r.Static("/static", "./public")

	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("__usid", store))

	ginext.LoadTemplates(r, "tpl")
	routes.SetRouters(r)
	r.Run(":8000")
}



//
//func weatherHandler(c *gin.Context) {
//	r := c.Request
//	begin := time.Now()
//	city := c.Params.ByName("city")
//
//	pt := r.URL.Query().Get("pt")
//	log.Printf("passed query parameter pt :: %v", pt)
//
//	var m weatherProvider
//
//	ovm := openWeatherMap{}
//	wug := weatherUnderGround{apiKey: "c5e7c706ec76acd6"}
//
//	if pt == "s" {
//		m = serialMetaWeatherProvider{ovm, wug}
//	} else {
//		// default means parallel fetching of weather info
//		m = parallelMetaWeatherProvider{ovm, wug}
//	}
//	data, err := m.temperature(city)
//	if err != nil {
//		c.Fail(http.StatusInternalServerError, err)
//		return
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"name": city,
//		"temp": data,
//		"took": time.Since(begin).String(),
//	})
//}
