package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"honeypot/admin/tools"
	mylog "honeypot/utils/log"
	"html/template"
	"io/fs"
	"net/http"
)

var Adminurl = tools.RandomAdminUrl()

func Admin() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(mylog.HttpLog("admin", "log"))
	fe, _ := fs.Sub(StaFS, "jkxtmw")
	router.StaticFS("jkxtmw", http.FS(fe))
	router.StaticFS(Adminurl+"/jkxtmw", http.FS(fe))
	templ := template.Must(template.New("").ParseFS(HtmlFS, "html/*.html"))
	router.SetHTMLTemplate(templ)
	store := cookie.NewStore([]byte("secure"))
	//设置session过期时间为两小时
	store.Options(sessions.Options{MaxAge: 120 * 60})
	router.Use(sessions.Sessions("sessionid", store))
	router.GET("/", tools.IndexPage)
	{
		v1 := router.Group(Adminurl)
		{
			v1.GET("/", Pot)
			v1.GET("/login", Pot)
			v1.POST("/login", Login)
			v1.GET("/logout", logout)
			v1.GET("/log", log)
			v1.GET("/log/:logname", log)
			v1.POST("/log", log)
			v1.POST("/logcount", CountLog)
			v1.POST("/log/logcount", CountLog)
			v1.POST("/logsearch", LogSearch)
			v1.POST("/log/logsearch", LogSearch)
			v1.GET("/changepass", ChangePass)
			v1.POST("/changepass", ChangePass)
			v1.GET("/potconfig", PotConfig)
			v1.POST("/potconfig", PotConfig)
			v1.POST("/uploadfile", UploadFile)
			v1.POST("/potcontrol", PotControl)
		}
	}
	tools.PageNotFound(router)
	return router
}
