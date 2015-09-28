// Copyright 2013-2015 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dm

import (
//	"fmt"
//	"os"
	as "github.com/aerospike/aerospike-client-go"
	"github.com/gin-gonic/gin"
	"net/http"
//	ext "cludezhu/trygo/libs/ginext"
//	"time"
	"strconv"
)

var (
//	host      string = "192.168.1.29"
//	port      int    = 3000
	client	*as.Client
)

const (
	NS = "camel"
	DE_UPDATER = "de_ctrl"
	DE_UPDATER_RATE = "de_ctrl_rate"
)

/**
	below are public functions
 */
func ConnectAS(pHost string, pPort int){
	var err error
	client, err = as.NewClient(pHost, pPort)
	panicOnError(err)
}

/**
	below are router functions
 */
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "dm/index.html", gin.H{})
}

func GetRecord(c *gin.Context) {
	pNs := c.DefaultQuery("namespace", NS)
	pSet := c.DefaultQuery("set", DE_UPDATER)
	pKey := c.DefaultQuery("key", "undefined")

	rec := loadASRecord(pNs, pSet, pKey)

	rst := parseRecord(rec)
	if(rst != nil) {
		renderJson(c, true, rst )
	}else{
		renderJson(c, false, "record not found" )
	}

}


//func GetServerStatus(c *gin.Context) {
//	pNs := c.DefaultQuery("namespace", NS)
//	pSet := c.DefaultQuery("set", DE_UPDATER)
//	rec := loadASRecord(pNs, pSet, "servers")
//
//	if rec != nil {
//		servers := rec.Bins["bin"].(map[interface{}]interface{})
//		renderJson(c, true, toJsonMap(servers))
//	} else {
//		renderJson(c, false, "record not found" )
//	}
//}

//func GetUpdater(c *gin.Context) {
//	pNs := c.DefaultQuery("namespace", NS)
//	pSet := c.DefaultQuery("set", DE_UPDATER)
//	pId := c.DefaultQuery("id", "0")
//	pType := c.DefaultQuery("type", "s")
//	pDate := c.DefaultQuery("date", ext.Date(time.Now(), "Ymd"))
//	pKey := fmt.Sprintf("%s_ctrl_%s_%s", pType, pDate, pId)
//
//	rec := loadASRecord(pNs, pSet, pKey)
//
//	if rec != nil {
//		rst := rec.Bins["bin"].(map[interface{}]interface{})
//		renderJson(c, true, toJsonMap(rst))
//	} else {
//		renderJson(c, false, "record not found" )
//	}
//}


/**
	below are private functions
 */
func renderJson(c *gin.Context, status bool, data interface{}){
	if(status){
		c.JSON(http.StatusOK, gin.H{"status": 1, "data":data} )
	}else{
		c.JSON(http.StatusOK, gin.H{"status": 0, "msg":data})
	}
}

func loadASRecord(ns string, set string, pkey string) (*as.Record) {
	var err error
	key, err := as.NewKey(ns, set, pkey)
	panicOnError(err)

	policy := as.NewPolicy()
	rec, err := client.Get(policy, key)
	panicOnError(err)

	return rec
}

func parseRecord(rec *as.Record) (gin.H){
	if rec != nil {
		rst := gin.H{}
		for k, v := range rec.Bins {
			rst[k] = toJson(v)
		}
		return rst
	} else {
		return nil
	}
}

func toJson(source interface{}) (interface{}) {
	switch v := source.(type) {
	case []interface{}:
		rst := []interface{}{}
		for _, item := range v {
			rst = append(rst, toJson(item))
		}
		return rst

	case map[interface{}]interface{}:
		rst := gin.H{}
		for ik, iv := range v {
			rst[toStr(ik)] = toJson(iv)
		}
		return rst

	default:
		return v
	}
}

func toStr(arg interface{}) string {
	var rst string
	switch v := arg.(type) {
	case int:
		rst = strconv.Itoa(v)
	case string:
		rst = v
	default:
		rst = ""
	}
	return rst
}

func toJsonMap(vmap map[interface{}]interface{}) (gin.H){
	rst := gin.H{}
	for k, v := range vmap {
		switch v.(type) {
			case map[interface{}]interface{}:
				rst[k.(string)] = toJsonMap(v.(map[interface{}]interface{}))
			default:
				rst[k.(string)] = v
		}
	}
	return rst
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
