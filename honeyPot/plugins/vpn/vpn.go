package vpn

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

func VPN() http.Handler {
	router := gin.New()
	router.Use(mylog.HttpLog("vpn", "vpnlog"))
	router.Use(static.Serve("/", plugins.EmbedFolder(resource, "resource")))
	router.StaticFile("/upload/EasyConnectInstaller.exe", tools.VpnFile)
	return router
}
