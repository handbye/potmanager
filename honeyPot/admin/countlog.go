package admin

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"honeypot/admin/tools"
	"net/http"
)

func CountLog(c *gin.Context) {
	islogin := tools.GetSession(c)
	starttime := c.PostForm("startDate")
	endtime := c.PostForm("endDate")
	ip := c.PostForm("ip")
	method := c.PostForm("reqmethod")
	tablename := c.PostForm("tablename")
	if tablename == ""{
		tablename = "log"
	}
	var datanum int
	if tools.In(tablename, tools.Config("httplog")) {
		datanum = QueryHttpCountLog(tablename, starttime, endtime, ip, method)
	}
	if tools.In(tablename, tools.Config("nohttplog")) {
		datanum = QueryNoHttpCountLog(tablename, starttime, endtime)
	}
	if islogin {
		if c.Request.Method == "POST" {
			c.JSON(http.StatusOK, gin.H{
				"datanum": datanum,
			})
		}
	}
}

func QueryHttpCountLog(tablename string, starttime string, endtime string, ip string, method string) (datanum int) {
	tools.SafeDate(tablename)
	tools.SafeDate(starttime)
	tools.SafeDate(endtime)
	tools.SafeDate(ip)
	tools.SafeDate(method)
	db, err := sql.Open("sqlite3", tools.DbPath)
	checkErr(err)
	// 查询数据
	var querystring string
	var num int
	if tools.In(tablename, tools.Config("httplog")) {
		if starttime == "" && endtime == "" && ip == "" && method == "" {
			querystring = fmt.Sprintf("SELECT COUNT(1) FROM %s", tablename)
		}
		if starttime == "" && endtime == "" && ip != "" && method == "" {
			querystring = fmt.Sprintf("SELECT COUNT(1) FROM %s  WHERE clientIP='%s'", tablename, ip)
		}
		if starttime == "" && endtime == "" && ip == "" && method != "" {
			querystring = fmt.Sprintf("SELECT COUNT(1) FROM %s  WHERE reqMethod='%s'", tablename, method)
		}
		if starttime == "" && endtime == "" && ip != "" && method != "" {
			querystring = fmt.Sprintf("SELECT COUNT(1) FROM %s  WHERE clientIP='%s' and reqMethod='%s'", tablename, ip, method)
		}
		if starttime != "" && endtime != "" && ip == "" && method == "" {
			querystring = fmt.Sprintf("SELECT COUNT(1) FROM %s  WHERE time >='%s' AND Time <='%s'", tablename, starttime, endtime)
		}
		if starttime != "" && endtime != "" && ip != "" && method != "" {
			querystring = fmt.Sprintf("SELECT COUNT(1) FROM %s  WHERE time >='%s' AND Time <='%s' AND clientIP='%s' and reqMethod='%s'", tablename, starttime, endtime, ip, method)
		}
		if starttime != "" && endtime != "" && ip == "" && method != "" {
			querystring = fmt.Sprintf("SELECT COUNT(1) FROM %s  WHERE time >='%s' AND Time <='%s' AND reqMethod='%s", tablename, starttime, endtime, method)
		}
		if starttime != "" && endtime != "" && ip != "" && method == "" {
			querystring = fmt.Sprintf("SELECT COUNT(1) FROM %s  WHERE time >='%s' AND Time <='%s' AND clientIP='%s'", tablename, starttime, endtime, ip)
		}
		rows, err := db.Query(querystring)
		if err != nil {
			fmt.Printf("数据库查询数据失败：%s\n", err)
		} else {
			for rows.Next() {
				err = rows.Scan(&num)
				if err == nil {
					datanum = num
				} else {
					datanum = 1
				}
			}
			db.Close()
		}
	}
	return datanum
}

func QueryNoHttpCountLog(tablename string, starttime string, endtime string) (datanum int) {
	tools.SafeDate(tablename)
	tools.SafeDate(starttime)
	tools.SafeDate(endtime)
	db, err := sql.Open("sqlite3", tools.DbPath)
	checkErr(err)
	// 查询数据
	var querystring string
	var num int
	if tools.In(tablename, tools.Config("nohttplog")) {
		if starttime == "" && endtime == "" {
			querystring = fmt.Sprintf("SELECT COUNT(1) FROM %s", tablename)
		}
		if starttime != "" && endtime != "" {
			querystring = fmt.Sprintf("SELECT COUNT(1) FROM %s  WHERE time >='%s' AND Time <='%s'", tablename, starttime, endtime)
		}
		rows, err := db.Query(querystring)
		if err != nil {
			fmt.Printf("数据库查询数据失败：%s\n", err)
		} else {
			for rows.Next() {
				err = rows.Scan(&num)
				if err == nil {
					datanum = num
				} else {
					datanum = 1
				}
			}
			db.Close()
		}
	}
	return datanum
}
