package admin

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"honeypot/admin/tools"
	"net/http"
)

type user struct {
	username string
	password string
}

var State = make(map[string]interface{})

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func queryData() (l []user, e error) {
	db, err := sql.Open("sqlite3", tools.DbPath)
	checkErr(err)
	// 查询数据
	rows, err := db.Query("SELECT USER,PASS FROM USER WHERE ID =1")
	checkErr(err)
	var result = make([]user, 0)
	for rows.Next() {
		var username,password string
		err = rows.Scan(&username, &password)
		checkErr(err)
		result = append(result, user{username, password})
	}
	db.Close()
	return result, nil
}

// Login 登录
func Login(c *gin.Context) {
	if c.Request.Method == "POST"{
		username := c.PostForm("username")
		password := c.PostForm("password")
		passmd5 :=  md5.Sum([]byte(password))
		passstr := fmt.Sprintf("%x", passmd5)
		res, _ := queryData()
		if username == res[0].username && passstr == res[0].password{
			State["state"]=1
			State["text"]="登录成功"
			session := sessions.Default(c)
			session.Set("secure", "admin")
			session.Save()
			c.JSON(http.StatusOK, gin.H{"code": 1, "message": "登录成功"})
		}else{
			State["state"]=0
			State["text"]="用户名或密码错误"
			c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
			//c.HTML(http.StatusOK, "login.html", honeyPot.H{
			//	"result": State["text"],
			//})
		}
	}
}
