package admin

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"honeypot/admin/tools"
	"net/http"
)

func updateData(pass string) bool {
	db, err := sql.Open("sqlite3", tools.DbPath)
	checkErr(err)
	// 更新数据
	querystring := fmt.Sprintf("update USER SET PASS = '%s' WHERE id =1", pass)
	_, err = db.Exec(querystring)
	if err != nil {
		return false
	}
	return true
}

func ChangePass(c *gin.Context) {
	islogin := tools.GetSession(c)
	if islogin {
		if c.Request.Method == "GET" {
			c.HTML(http.StatusOK, "pwd.html", gin.H{
				"IsLogin": islogin,
			})
		}
		if c.Request.Method == "POST" {
			oldpass := c.PostForm("oldpass")
			password1 := c.PostForm("password1")
			password2 := c.PostForm("password2")
			passmd5 := md5.Sum([]byte(oldpass))
			passstr := fmt.Sprintf("%x", passmd5)
			res, _ := queryData()
			if oldpass != "" && password1 != "" && password2 != "" {
				if passstr != res[0].password {
					c.JSON(http.StatusOK, gin.H{"code": -1, "message": "旧密码错误"})
					return
				}
				if password1 != password2 {
					c.JSON(http.StatusOK, gin.H{"code": 0, "message": "两次输入的密码不一致"})
					return
				}
				if !tools.CheckPass(8, 20, 4, password1) && !tools.CheckPass(8, 15, 4, password1) {
					c.JSON(http.StatusOK, gin.H{"code": 3, "message": "密码长度必须大于8位，并且必须包含大小写字母,数字和特殊符号"})
				} else {
					passmd51 := md5.Sum([]byte(password1))
					passstr1 := fmt.Sprintf("%x", passmd51)
					result := updateData(passstr1)
					if result {
						c.JSON(http.StatusOK, gin.H{"code": 1, "message": "密码修改成功"})
					} else {
						c.JSON(http.StatusOK, gin.H{"code": 2, "message": "密码修改失败"})
						return
					}
				}
			}
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}
