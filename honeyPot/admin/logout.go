package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func logout(c *gin.Context)  {

	//清除该用户登录状态的数据
	session := sessions.Default(c)
	session.Delete("secure")
	session.Save()
	//session.Clear()

	c.Redirect(302,"/"+Adminurl)
}