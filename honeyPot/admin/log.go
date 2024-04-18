package admin

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"honeypot/admin/tools"
	"net/http"
	"strconv"
)

type HttpLogData struct {
	Time         string
	ClientIP     string
	StatusCode   string
	ReqMethod    string
	ReqUri       string
	Full_message string
}
type NohttpLogData struct {
	Time         string
	Full_message string
}

func log(c *gin.Context) {
	islogin := tools.GetSession(c)
	page := c.PostForm("page")
	var i int
	i, _ = strconv.Atoi(page)
	if i == 0 {
		i = 1
	}
	starttime := c.PostForm("startDate")
	endtime := c.PostForm("endDate")
	ip := c.PostForm("ip")
	method := c.PostForm("reqmethod")
	var tablename string
	if c.Param("logname") != ""{
		tablename = c.Param("logname")
	}else{
		tablename = "log"
	}
	if tools.In(tablename, tools.Config("httplog")){
		res, err := QuerHttpyLog(tablename,i, starttime, endtime, ip, method)
		if islogin {
			if err != nil || len(res) == 0 {
				res = append(res, HttpLogData{"", "", "", "", "", ""})
			}
			if len(res) != 0 {
				if c.Request.Method == "GET" {
					c.HTML(http.StatusOK, "log.html", gin.H{
						"adminurl": Adminurl,
						"IsLogin":  islogin,
						"res":      res,
					})
				}

				if c.Request.Method == "POST" {
					c.JSON(http.StatusOK, gin.H{
						"res": res,
					})
				}
			}
		} else {
			c.HTML(http.StatusOK, "login.html", nil)
		}
	}
	if tools.In(tablename, tools.Config("nohttplog")){
		res, err := QuerNoHttpyLog(tablename,i, starttime, endtime)
		if islogin {
			if err != nil || len(res) == 0 {
				res = append(res, NohttpLogData{"", ""})
			}
			if len(res) != 0 {
				if c.Request.Method == "GET" {
					c.HTML(http.StatusOK, "log1.html", gin.H{
						"adminurl": Adminurl,
						"IsLogin":  islogin,
						"res":      res,
					})
				}

				if c.Request.Method == "POST" {
					c.JSON(http.StatusOK, gin.H{
						"res": res,
					})
				}
			}
		} else {
			c.HTML(http.StatusOK, "login.html", nil)
		}
	}
}

func QuerHttpyLog(tablename string, page int, starttime string, endtime string, ip string, method string) (data []HttpLogData, e error) {
	tools.SafeDate(tablename)
	tools.SafeDate(starttime)
	tools.SafeDate(endtime)
	tools.SafeDate(ip)
	tools.SafeDate(method)
	db, err := sql.Open("sqlite3", tools.DbPath)
	checkErr(err)
	// 查询数据
	var querystring string
	var result = make([]HttpLogData, 0)
	if tools.In(tablename, tools.Config("httplog")) {
		if starttime == "" && endtime == "" && ip == "" && method == "" {
			querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s ORDER BY id DESC LIMIT %d,%d", tablename, (page-1)*10, 10)
		}
		if starttime == "" && endtime == "" && ip != "" && method == "" {
			querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s  WHERE clientIP='%s' ORDER BY id DESC LIMIT %d,%d", tablename, ip, (page-1)*10, 10)
		}
		if starttime == "" && endtime == "" && ip == "" && method != "" {
			querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s  WHERE reqMethod='%s' ORDER BY id DESC LIMIT %d,%d", tablename, method, (page-1)*10, 10)
		}
		if starttime == "" && endtime == "" && ip != "" && method != "" {
			querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s  WHERE clientIP='%s' and reqMethod='%s' ORDER BY id DESC LIMIT %d,%d", tablename, ip, method, (page-1)*10, 10)
		}
		if starttime != "" && endtime != "" && ip == "" && method == "" {
			querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s  WHERE time >='%s' AND Time <='%s' ORDER BY id DESC  LIMIT %d,%d", tablename, starttime, endtime, (page-1)*10, 10)
		}
		if starttime != "" && endtime != "" && ip != "" && method != "" {
			querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s  WHERE time >='%s' AND Time <='%s' AND clientIP='%s' and reqMethod='%s' ORDER BY id DESC LIMIT %d,%d", tablename, starttime, endtime, ip, method, (page-1)*10, 10)
		}
		if starttime != "" && endtime != "" && ip == "" && method != "" {
			querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s  WHERE time >='%s' AND Time <='%s' AND reqMethod='%s' ORDER BY id DESC LIMIT %d,%d", tablename, starttime, endtime, method, (page-1)*10, 10)
		}
		if starttime != "" && endtime != "" && ip != "" && method == "" {
			querystring = fmt.Sprintf("SELECT time,clientIp,statusCode,reqMethod,reqUri,full_message FROM %s  WHERE time >='%s' AND Time <='%s' AND clientIP='%s' ORDER BY id DESC LIMIT %d,%d", tablename, starttime, endtime, ip, (page-1)*10, 10)
		}
		rows, err := db.Query(querystring)
		if err != nil {
			fmt.Printf("数据库查询数据失败：%s\n", err)
		} else {
			for rows.Next() {
				var (
					time, clientIp                  string
					statusCode                      string
					reqMethod, reqUri, full_message string
				)
				err = rows.Scan(&time, &clientIp, &statusCode, &reqMethod, &reqUri, &full_message)
				if err == nil {
					result = append(result, HttpLogData{time, clientIp, statusCode, reqMethod, reqUri, full_message})
				} else {
					result = append(result, HttpLogData{"", "", "", "", "", ""})
				}
			}
			defer db.Close()
		}
	}
	return result, err
}

func QuerNoHttpyLog(tablename string, page int, starttime string, endtime string) (data []NohttpLogData, e error) {
	tools.SafeDate(tablename)
	tools.SafeDate(starttime)
	tools.SafeDate(endtime)
	db, err := sql.Open("sqlite3", tools.DbPath)
	checkErr(err)
	// 查询数据
	var querystring string
	var result = make([]NohttpLogData, 0)
	if tools.In(tablename, tools.Config("nohttplog")) {
		if starttime == "" && endtime == "" {
			querystring = fmt.Sprintf("SELECT time,msg FROM %s ORDER BY id DESC LIMIT %d,%d", tablename, (page-1)*10, 10)
		}
		if starttime != "" && endtime != "" {
			querystring = fmt.Sprintf("SELECT time,msg FROM %s  WHERE time >='%s' AND Time <='%s' ORDER BY id DESC  LIMIT %d,%d", tablename, starttime, endtime, (page-1)*10, 10)
		}
		rows, err := db.Query(querystring)
		if err != nil {
			fmt.Printf("数据库查询数据失败：%s\n", err)
		} else {
			for rows.Next() {
				var (
					time         string
					full_message string
				)
				err = rows.Scan(&time, &full_message)
				if err == nil {
					result = append(result, NohttpLogData{time, full_message})
				} else {
					result = append(result, NohttpLogData{"", ""})
				}
			}
			defer db.Close()
		}
	}
	return result, err
}
