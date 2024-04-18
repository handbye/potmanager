package burpsuite

import (
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"honeypot/admin/tools"
	"honeypot/plugins"
	mylog "honeypot/utils/log"
	"net/http"
)

//go:embed resource
var resource embed.FS

func BurpSuite() http.Handler {
	router := gin.New()
	router.Use(mylog.HttpLog("BurpSuite", "burplog"))
	router.Use(static.Serve("/", plugins.EmbedFolder(resource, "resource")))
	router.StaticFile("/upload/api.js", tools.BurpFile)
	return router
}
