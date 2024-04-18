package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"honeypot/admin/tools"
	"net/http"
	"os"
)

func UploadFile(c *gin.Context) {
	islogin := tools.GetSession(c)
	if islogin {
		if c.Request.Method == "POST" {
			f, err := c.FormFile("file")
			code := 1
			if err == nil {
				//调用save之前先删除对应文件
				err = os.Remove(tools.VpnFile)
				if err != nil {
					fmt.Println(err)
				}

				//后续其他upload调用需修改dst参数
				err = c.SaveUploadedFile(f, tools.VpnFile)
				if err != nil {
					fmt.Println(err)
				} else {
					code = 0
				}
			}
			c.JSON(http.StatusOK, gin.H{"code": code})
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}
